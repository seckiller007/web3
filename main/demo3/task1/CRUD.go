package main

import (
	"fmt"
	"gorm.io/gorm"
)

func createStudents(db *gorm.DB) {
	students := []*Student{
		{Name: "Jinzhu", Age: 18, Grade: "5"},
		{Name: "Jackson", Age: 19, Grade: "3"},
	}

	result := db.Create(&students)
	if result.Error != nil {
		fmt.Println("创建学生记录失败:", result.Error)
		return
	}

	fmt.Println("创建成功，影响行数:", result.RowsAffected)
	for _, student := range students {
		fmt.Printf("创建的学生: ID=%d, 姓名=%s, 年龄=%d\n", student.ID, student.Name, student.Age)
	}
}
