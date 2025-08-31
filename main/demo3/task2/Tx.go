package main

import (
	"errors" // 自定义错误
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause" // GORM 的 SQL 子句，这里用来加行级锁
)

/************ 事务核心 ************/
// transfer 在一个事务中完成：余额检查 -> 扣款 -> 加款 -> 写交易流水
func transfer(db *gorm.DB, fromID, toID uint64, amount int64) error {
	// 前置校验 自转和非正金额直接拒绝。
	if fromID == toID {
		return errors.New("cannot transfer to self")
	}
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	// 开始事务
	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("begin transaction failed: %w", tx.Error)
	}

	// 延迟处理：如果发生错误则回滚事务
	defer func() {
		if r := recover(); r != nil {
			// 处理panic情况
			tx.Rollback()
		} else if tx.Error != nil {
			// 处理普通错误
			tx.Rollback()
		}
	}()

	// 锁定转出账户并查询（悲观锁）
	//防止余额超扣：
	//假设用户 A（余额 100 元）同时向用户 B 转账 50 元、向用户 C 转账 60 元。如果不锁定账户，可能出现两次查询都读到初始余额 100 元，最终导致 A 的余额变为100-50=50再减去 60，出现-10的负数（超扣）。
	//锁定后，第一个事务会锁定 A 的记录，第二个事务必须等待第一个事务完成，才能读取 A 的最新余额（50 元），此时会发现余额不足而拒绝第二次转账，避免超扣。
	//保证余额校验有效：避免在扣减的时候余额不足，被另一个事务转账转走了
	//代码中先查询转出账户余额（fromAccount.Balance < amount），再执行扣减。如果不锁定，两次操作之间可能有其他事务修改了余额，导致校验结果失效（例如校验时余额充足，但扣减时余额已被其他事务转走）。
	//锁定后，其他事务无法修改该账户，确保校验和扣减的原子性。
	var fromAccount Account
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", fromID).
		First(&fromAccount).Error; err != nil {
		return fmt.Errorf("failed to get from account: %w", err)
	}

	// 检查余额是否充足
	if fromAccount.Balance < amount {
		return errors.New("insufficient balance")
	}

	// 锁定转入账户（悲观锁）
	var toAccount Account
	//防止数据更新丢失：
	//假设用户 B 同时接收用户 A 的 50 元和用户 C 的 30 元转账。如果不锁定 B 的账户，两个事务可能同时读取 B 的初始余额（例如 200 元），分别计算为200+50=250和200+30=230，最终覆盖写入时，
	//后提交的事务会覆盖先提交的结果（导致其中一笔转账金额「丢失」）。
	//锁定后，两个事务会排队执行，先完成的事务将 B 的余额更新为 250 元，后执行的事务会基于 250 元计算（250+30=280），确保两笔金额都被正确累加。
	//避免并发更新冲突：
	//即使转入账户不涉及余额不足的校验，也需要锁定。因为数据库的「写操作」本身可能触发冲突（例如两个事务同时修改同一行记录），锁定可以确保更新按顺序执行，避免数据库抛出并发更新错误（如 MySQL 的1062错误）。
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", toID).
		First(&toAccount).Error; err != nil {
		return fmt.Errorf("failed to get to account: %w", err)
	}

	// 扣减转出账户余额
	if err := tx.Model(&Account{}).
		Where("id = ?", fromID).
		Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		return fmt.Errorf("failed to update from account: %w", err)
	}

	// 增加转入账户余额
	if err := tx.Model(&Account{}).
		Where("id = ?", toID).
		Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		return fmt.Errorf("failed to update to account: %w", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit transaction failed: %w", err)
	}

	return nil

}
func initDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/web3?parseTime=true&charset=utf8mb4&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("failed to connect database: " + err.Error())
		return nil
	}
	return db
}

func main() {
	// 初始化数据库连接
	db := initDB()
	if db == nil {
		fmt.Println("数据库连接失败，程序退出")
		return
	}

	transfer(db, 1, 2, 5000) //执行完观察数据库数据变化

	// 执行数据操作
	//createStudents(db)

}
