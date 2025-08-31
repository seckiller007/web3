package main

import (
	"fmt"
)

func main() {
	// 初始化数据库连接
	db := initDB()
	if db == nil {
		fmt.Println("数据库连接失败，程序退出")
		return
	}

	// 自动迁移数据表
	migrateTables(db)

	// 执行数据操作
	//createStudents(db)

}
