/**
@author ZhengHao Lou
Date    2022/04/30
*/
package main

import "math"

/**
https://leetcode-cn.com/problems/smallest-range-i/
*/
func smallestRangeI(nums []int, k int) int {
	var mi, mx = math.MaxInt32, math.MinInt32
	for i := range nums {
		mx = max(mx, nums[i])
		mi = min(mi, nums[i])
	}
	var s = mx - mi
	return min(s, max(s-k<<1, 0))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
