package main

import "fmt"

func plusOne(digits []int) []int {
	//逆序遍历
	for i := len(digits) - 1; i >= 0; i-- {
		//加一也就是个位+1，最后位+1，若末位加一进一，取模会等于0，继续循环
		digits[i]++
		digits[i] %= 10 //加完之后对10取模，若为0说明这一位是9
		if digits[i] != 0 {
			return digits
		}
	}
	//循环结束所有位加完仍然取模是0，特殊全是9的情况
	return append([]int{1}, digits...)
}

func main() {
	//fmt.Println(plusOne([]int{1, 2, 3, 9})) // [1 2 4]
	fmt.Println(plusOne([]int{9, 9, 9})) // [1 0 0 0]
}
