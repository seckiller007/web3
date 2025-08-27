package main

import "fmt"

func removeDuplicates(nums []int) int {
	var i = 0
	for j := 1; j < len(nums); j++ {
		//慢指针对应的数和快指针对应的数不相等，移动慢指针，相等继续遍历快指针，慢指针不动，最后得到慢指针指到的位置就是不重复元素
		if nums[i] != nums[j] {
			nums[i+1] = nums[j] //用下一个不等的元素替换每次首次出现相等元素的位置
			i++
		}
	}
	return len(nums[:i+1])
}

func main() {
	//fmt.Println(plusOne([]int{1, 2, 3, 9})) // [1 2 4]
	fmt.Println(removeDuplicates([]int{1, 2, 3, 3, 9})) // [1 0 0 0]
}
