package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64      // 面积
	Perimeter() float64 //周长
}

// 矩形
type Rectangle struct {
	Width  float64 // 宽
	Height float64 // 高
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 圆形
type Circle struct {
	Radius float64 // 半径
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Println("Employee:", e.Person.Name, e.Person.Age, e.EmployeeID)
}

func ObjectExercise() {
	fmt.Println("-------面向对象练习-------")
	/*
		题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
		考察点 ：接口的定义与实现、面向对象编程风格。
	*/
	//InterfaceExercise()

	/*
		题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
		考察点 ：组合的使用、方法接收者。
	*/
	StructCombinationExercise()
}

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
func InterfaceExercise() {
	fmt.Println("-----接口使用-----")
	// 矩形
	rectangle := &Rectangle{Width: 10, Height: 5}
	fmt.Printf("长: %.2f, 宽: %.2f 面积: %.2f 周长: %.2f\n", rectangle.Width, rectangle.Height, rectangle.Area(), rectangle.Perimeter())
	// 圆形
	circle := &Circle{Radius: 5}
	fmt.Printf("半径: %.2f 面积: %.2f 周长: %.2f\n", circle.Radius, circle.Area(), circle.Perimeter())
}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
func StructCombinationExercise() {
	fmt.Println("-----组合使用-----")
	// 创建一个 Person 实例
	person := Person{Name: "张三", Age: 18}
	// 创建一个 Employee 实例，并组合 Person 实例
	employee := Employee{Person: person, EmployeeID: 1001}
	// 调用 PrintInfo() 方法输出员工信息
	employee.PrintInfo()
}
