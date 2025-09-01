package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func initDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %v", err)
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	return db, nil
}

type Book struct {
	ID     int     `db:"id"`     // 书籍ID
	Title  string  `db:"title"`  // 书籍标题
	Author string  `db:"author"` // 作者
	Price  float64 `db:"price"`  // 价格
}

func main() {
	// 数据库连接信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/web3?parseTime=true&charset=utf8mb4&loc=Local"
	// 初始化数据库连接
	db, err := initDB(dsn)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer db.Close()

	// 查询价格大于50元的书籍
	var books []Book
	query := "SELECT id, title, author, price FROM books WHERE price > ?"
	err = db.Select(&books, query, 50.0)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	// 打印查询结果
	fmt.Printf("找到 %d 本价格大于50元的书籍：\n", len(books))
	for i, book := range books {
		fmt.Printf("%d. %s (作者: %s) - 价格: %.2f元\n",
			i+1, book.Title, book.Author, book.Price)
	}
}
