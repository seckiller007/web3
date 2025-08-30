package main

import (
	"fmt"
)

// 指针  值传递不会修改副本的值，引用传递会修改
func Pointer() {
	fmt.Println("--------指针--------")
	// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	// 考察点 ：指针的使用、值传递与引用传递的区别。
	var a int = 10
	fmt.Printf("a地址= = %s, a = %d\n", &a, a)
	add(&a)
	fmt.Printf("a地址= = %s, a = %d\n", &a, a)

	// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	// 考察点 ：指针运算、切片操作。
	array := []int{1, 2, 3, 4, 5}
	fmt.Printf("before a = %v\n", array)
	multiplication(&array)
	fmt.Printf("after a = %v\n", array)
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func multiplication(array *[]int) {
	for i, v := range *array {
		(*array)[i] = v * 2
	}
}

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func add(a *int) {
	*a += 10
}

func add2(a int) {
	a = 100
}
