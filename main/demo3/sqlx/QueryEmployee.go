package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL驱动
	"github.com/jmoiron/sqlx"
)

// 查询所有技术部员工
func getTechDepartmentEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	// 使用Select查询多条记录
	err := db.Select(&employees, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	return employees, nil
}

// 查询工资最高的员工
func getHighestSalaryEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	// 使用Get查询单条记录
	err := db.Get(&employee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		return Employee{}, fmt.Errorf("查询失败: %v", err)
	}
	return employee, nil
}

func main() {
	// 数据库连接信息 (根据实际情况修改)
	dsn := "root:123456@tcp(127.0.0.1:3306)/web3?parseTime=true&charset=utf8mb4&loc=Local"

	// 初始化数据库连接
	db, err := initDB(dsn)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer db.Close()

	// 查询技术部员工
	techEmployees, err := getTechDepartmentEmployees(db)
	if err != nil {
		log.Printf("查询技术部员工失败: %v", err)
	} else {
		fmt.Println("技术部员工:")
		for _, emp := range techEmployees {
			fmt.Printf("ID: %d, 姓名: %s, 工资: %.2f\n", emp.ID, emp.Name, emp.Salary)
		}
	}

	// 查询工资最高的员工
	topEmployee, err := getHighestSalaryEmployee(db)
	if err != nil {
		log.Printf("查询工资最高员工失败: %v", err)
	} else {
		fmt.Printf("\n工资最高的员工: ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
			topEmployee.ID, topEmployee.Name, topEmployee.Department, topEmployee.Salary)
	}
}
