/*
*
@author ZhengHao Lou
Date    2022/12/05
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/minimum-score-of-a-path-between-two-cities/
*/

type pair struct {
	x, y int
}

func minScore(n int, roads [][]int) int {
	g := make([][]pair, n+1)
	for _, r := range roads {
		g[r[0]] = append(g[r[0]], pair{r[1], r[2]})
		g[r[1]] = append(g[r[1]], pair{r[0], r[2]})
	}

	var vis = make([]bool, n+1)
	var dfs func(i, pi int)
	dfs = func(i, pi int) {
		for _, np := range g[i] {
			ni, _ := np.x, np.y
			if vis[ni] {
				continue
			}
			vis[ni] = true
			dfs(ni, i)
		}
	}

	var mi = math.MaxInt32
	vis[1] = true
	dfs(1, 1)
	for i := 1; i <= n; i++ {
		if vis[i] {
			for _, np := range g[i] {
				mi = min(mi, np.y)
			}
		}
	}

	return mi
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}))
	fmt.Println(minScore(4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}))
	fmt.Println(minScore(6, [][]int{{4, 5, 7468}, {6, 2, 7173}, {6, 3, 8365}, {2, 3, 7674}, {5, 6, 7852}, {1, 2, 8547},
		{2, 4, 1885}, {2, 5, 5192}, {1, 3, 4065}, {1, 4, 7357}})) // 1885
}
