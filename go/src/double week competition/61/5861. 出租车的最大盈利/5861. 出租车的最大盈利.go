package main

import (
	"fmt"
	"sort"
)

/**
思路：记忆化搜索 + 二分
 */
func maxTaxiEarnings(n int, rides [][]int) int64 {
	m := len(rides)
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][0] <= rides[j][0]
	})

	fmt.Println(rides)

	next := make([]int, m)
	for i := 0; i < m; i++ {
		var j = lowerBound(rides, i + 1, rides[i][1])
		next[i] = j
	}

	mem := make(map[int]int64)
	var dfs func(int) int64
	dfs = func(i int) int64 {
		if c, ok := mem[i]; ok {
			return c
		}
		if i == m {
			return 0
		}

		mem[i] = max(dfs(next[i]) + int64(rides[i][1] - rides[i][0] + rides[i][2]), dfs(i + 1))
		return mem[i]
	}

	return dfs(0)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func lowerBound(ride [][]int, l, x int) int {
	r := len(ride)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if ride[m][0] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(maxTaxiEarnings(5, [][]int{{2,5,4},{1,5,1}}))
	fmt.Println(maxTaxiEarnings(5, [][]int{{1,6,1},{3,10,2},{10,12,3},{11,12,2},{12,15,2},{13,18,1}}))
	fmt.Println(maxTaxiEarnings(5, [][]int{{2,3,6},{8,9,8},{5,9,7},{8,9,1},{2,9,2},{9,10,6},{7,10,10},
		{6,7,9},{4,9,7},{2,3,1}}))	// 33
	fmt.Println(maxTaxiEarnings(5, [][]int{{9,10,2},{4,5,6},{6,8,1},{1,5,5},{4,9,5},{1,6,5},{4,8,3},{4,7,10},
		{1,9,8},{2,3,5}}))
}