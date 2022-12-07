/**
@author ZhengHao Lou
Date    2022/10/16
*/
package main

/**
https://leetcode.cn/problems/create-components-with-same-value/
思路：dfs
*/
func componentValue(nums []int, edges [][]int) int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	var target int
	var dfs func(i, pi int) int
	dfs = func(i, pi int) int {
		sum := nums[i]
		for _, ni := range g[i] {
			if ni != pi {
				res := dfs(ni, i)
				if res < 0 {
					return -1
				}
				sum += res
			}
		}

		if sum > target {
			return -1
		}
		if sum == target {
			return 0
		}
		return sum
	}

	var mx, s int
	for i := range nums {
		s += nums[i]
		mx = max(mx, nums[i])
	}

	for i := min(n, s/mx); ; i-- {
		if s%i == 0 {
			target = s / i
			if dfs(0, -1) == 0 {
				return i - 1
			}
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
