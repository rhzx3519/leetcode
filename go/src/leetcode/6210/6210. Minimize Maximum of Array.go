/**
@author ZhengHao Lou
Date    2022/10/16
*/
package main

import "sort"

/**
https://leetcode.cn/problems/create-components-with-same-value/
*/
func minimizeArrayValue(nums []int) int {
	var mx int
	for i := range nums {
		mx = max(mx, nums[i])
	}
	return sort.Search(mx, func(limit int) bool {
		var extra int
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
