package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	//按照左边界从小到大顺序排序  intervals[i][0] 第一个中括号是第几个区间，第二个[]是区间里的左右边界，[0]左边界
	sort.Slice(intervals, func(i, j int) bool {
		// intervals内部不停比较，按照左边界升序排列
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0] //不断更新前一个interval
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]   //当前区间
		if prev[1] < cur[0] { // 没有一点重合  上一个区间有边界小于当前的左边界，没有重合
			res = append(res, prev) //上一区间添加到结果集
			prev = cur              //更新上一区间
		} else { // 有重合 更新右边界，接着循环
			prev[1] = max(prev[1], cur[1])
		}
	}
	//容易遗漏最后一个prev，当所有的都遍历完之后剩下的那个prev
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // [1 0 0 0]
}
