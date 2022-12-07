/**
@author ZhengHao Lou
Date    2022/08/20
*/
package main

/**
https://leetcode.cn/problems/maximum-segment-sum-after-removals/
并查集
*/
func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	n := len(nums)
	sums := make([]int64, n+1)
	pa := make([]int, n+1)
	for i := range pa {
		pa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if pa[x] != x {
			pa[x] = find(pa[x])
		}
		return pa[x]
	}

	var ans = make([]int64, n)
	for i := n - 1; i > 0; i-- {
		x := removeQueries[i]
		to := find(x + 1)
		pa[x] = to // merge x, x+1
		sums[to] += sums[x] + int64(nums[x])
		ans[i-1] = max(ans[i], sums[to])
	}
	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
