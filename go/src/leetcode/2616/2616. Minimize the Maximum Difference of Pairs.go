package main

import "sort"

/*
*
https://leetcode.cn/problems/minimize-the-maximum-difference-of-pairs/
思路：二分 + 贪心
排序nums, 假设len(nums) >= 4, 选出两个数对
min(nums[1] - nums[0], nums[3] - nums[2]) 必然小于 min(nums[3] - nums[1], nums[2] - nums[0])
*/
func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	n := len(nums)
	return sort.Search(1e9, func(mx int) bool {
		var c int
		for i := 1; i < n; i++ {
			if nums[i]-nums[i-1] <= mx {
				c++
				i++
			}
		}
		return c >= p
	})
}
