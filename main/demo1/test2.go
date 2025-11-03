package main

/**
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reverse := 0
	//对x进行每个位上的数字获取
	for x > reverse { //x大于等于reverse时，循环结束了,只需要反转一半
		//不断更新反转数字
		reverse = reverse*10 + x%10
		x = x / 10 //x不断减除一位后的剩余数字
	}
	//reverse/10 对于x是奇数去除处于中位的数字。奇数的时候x<reverse,若x == reverse/10是回文数
	return x == reverse || x == reverse/10
}

func main() {
	x := 12121
	if isPalindrome(x) {
		fmt.Println(x, "是回文数")
	} else {
		fmt.Println(x, "不是回文数")
	}
}
