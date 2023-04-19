package main

import "sort"

/*
*
https://leetcode.cn/problems/widest-vertical-area-between-two-points-containing-no-points/
*/
func maxWidthOfVerticalArea(points [][]int) int {
	cnt := make(map[int]int)
	for i := range points {
		cnt[points[i][0]]++
	}
	var nums []int
	for x := range cnt {
		nums = append(nums, x)
	}
	sort.Ints(nums)
	n := len(nums)
	if n == 1 {
		return 0
	}
	var ans int
	for i := 0; i < n-1; i++ {
		ans = max(ans, nums[i+1]-nums[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
