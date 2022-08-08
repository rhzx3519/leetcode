/**
@author ZhengHao Lou
Date    2022/05/19
*/
package main

import "sort"

/**
https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/
*/
func minMoves2(nums []int) (c int) {
	n := len(nums)
	sort.Ints(nums)
	x := nums[n>>1]
	for _, num := range nums {
		c += abs(num - x)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}