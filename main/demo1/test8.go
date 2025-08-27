package main

func twoSum(nums []int, target int) []int {
	idx := map[int]int{}     // 创建一个空哈希表 map[键类型]值类型,{}表示空数组  make(map[int]int)
	for j, x := range nums { // 索引 j
		// 在左边找 nums[i]，满足 nums[i]+x=target
		if i, ok := idx[target-x]; ok { // 遍历到j的时候发现  nums[i]+x=target
			return []int{i, j} // 返回两个数的下标
		}
		idx[x] = j // 保存 nums[j] 和 j到map里
	}
	return nil // 题目保证一定有解，代码不会执行到这里
}
