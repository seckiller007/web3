package main

import (
	"fmt"
)

func findSingleNumber(nums []int) int {
	//定义一个map结构，key 是数字，value是数字出现的次数
	countMap := make(map[int]int)

	// 记录每个元素出现的次数 索引和对应的元素。如果你只需要元素而不需要索引，可以使用空白标识符 _ 来替代索引变量
	for _, num := range nums {
		countMap[num]++
	}

	// 找到出现次数为1的元素
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	// 如果没有找到只出现一次的元素，返回0（根据题目描述这种情况不会发生）
	return 0
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	singleNumber := findSingleNumber(nums)
	fmt.Println("只出现一次的数字是:", singleNumber)
}
