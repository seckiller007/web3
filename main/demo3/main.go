//package main
//
//import (
//	"fmt"
//)
//
//// 返回值一个函数 func (int) int    第一个int是参数，第二个int是函数返回值
//func getSum() func(int) int {
//	var sum int = 0
//	return func(num int) int {
//		sum = num + sum
//		return sum
//	}
//}
//
//// 一个可以返回多个值的函数
//func numbers() (int, int, string) {
//	a, b, c := 1, 2, "str"
//	return a, b, c
//}
//
//func main() {
//	//s2 := make([]int, 3, 4)
//	//fmt.Println(s2)
//	//
//	//s2 = append(s2, 7)
//	//fmt.Println(s2)
//	//
//	//s3 := make([]int, 4)
//	//fmt.Println(s3)
//	//
//	//var intarr [6]int = [6]int{1, 2, 3, 4, 5, 6}
//	////切出一片数组。从1到3，左闭右开
//	//var slice []int = intarr[1:3]
//	//fmt.Println(slice)
//	////_的用法
//	//_, numb, strs := numbers() //只获取函数返回值的后两个，_
//	//fmt.Println(numb, strs)
//	f := getSum()
//	fmt.Println(f(1))
//	//db := initDB()
//	//if db == nil {
//	//	fmt.Println("数据库连接失败，程序退出")
//	//	return
//	//}
//	//
//	//// 自动迁移数据表
//	//migrateTables(db)
//	//
//	//// 执行数据操作
//	//createStudents(db)
//	//dsn := "root:123456@tcp(127.0.0.1:3306)/web3?parseTime=true&charset=utf8mb4&loc=Local"
//	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	//if err != nil {
//	//	println("failed to connect database: " + err.Error())
//	//}
//	//
//	//// 示例：创建文章（会自动触发Post的AfterCreate钩子）
//	//newPost := Post{Title: "测试文章", Content: "这是一篇测试文章", UserID: 1}
//	//db.Create(&newPost)
//	//
//	//// 示例：删除评论（会自动触发Comment的AfterDelete钩子）
//	//db.Delete(&Comment{}, 1)
//}
