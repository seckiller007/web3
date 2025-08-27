package main

import "strings"

// GO
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0] //假设第一个字符串为基准
	for _, k := range strs {
		//从下一个字符串开始，如果基准不在下一个字符串里，将基准进行截取，直到截取后的字符串在下一个字符串，如果截取到最后也没有，直接返回“”
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
