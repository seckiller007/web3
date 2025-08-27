package main

import (
	"fmt"
)

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	//go 语言中创建map，key是右括号，value左括号
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{} //构建slice类型的stack
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			//先进来的是右括号或这个右括号和对应的左括号不匹配，直接结束
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1] //否则的话，说明匹配上了就出栈
		} else {
			stack = append(stack, s[i]) //左括号入栈
		}
	}
	return len(stack) == 0
}

func main() {
	isValid("([])")
	fmt.Println("Hello World")
}
