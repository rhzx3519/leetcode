package main

import (
	"sort"
)

/*
*
https://leetcode.cn/problems/count-ways-to-group-overlapping-ranges/
思路：所有相交的元素归于一个集合，假设共有m个互不相交的集合，每个集合都可以决定自己在
第一组或者第二组，那么共有2^m种情况
*/
const MOD int = 1e9 + 7

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] <= ranges[j][0]
	})
	ans, maxR := 2, ranges[0][1]
	for _, p := range ranges {
		if p[0] > maxR { // 产生了新的不相交的集合, m++
			ans = (ans << 1) % MOD
		}
		maxR = max(maxR, p[1])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	countWays([][]int{{5, 15}, {6, 10}})
}
