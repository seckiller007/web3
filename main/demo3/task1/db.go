package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库连接
func initDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/web3?parseTime=true&charset=utf8mb4&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("failed to connect database: " + err.Error())
		return nil
	}
	return db
}

// 迁移数据表
func migrateTables(db *gorm.DB) {
	// 迁移Student表
	err := db.AutoMigrate(&Student{})
	if err != nil {
		println("迁移数据表失败: " + err.Error())
	}

}
