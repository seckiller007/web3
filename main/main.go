package main

import (
	"fmt"
)

// 返回值一个函数 func (int) int    第一个int是参数，第二个int是函数返回值
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = num + sum
		return sum
	}
}

func main() {
	s2 := make([]int, 3, 4)
	fmt.Println(s2)

	s2 = append(s2, 7)
	fmt.Println(s2)

	s3 := make([]int, 4)
	fmt.Println(s3)

	var intarr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	//切出一片数组。从1到3，左闭右开
	var slice []int = intarr[1:3]
	fmt.Println(slice)

	f := getSum()
	fmt.Println(f(1))
}
