/**
@author ZhengHao Lou
Date    2022/02/12
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/number-of-enclaves
*/
var off = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		for _, dij := range off {
			ni, nj := dij[0]+i, dij[1]+j
			if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
				grid[ni][nj] = 0
				dfs(ni, nj)
			}
		}
	}

	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			grid[i][0] = 0
			dfs(i, 0)
		}
		if grid[i][n-1] == 1 {
			grid[i][n-1] = 0
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if grid[0][j] == 1 {
			grid[0][j] = 0
			dfs(0, j)
		}
		if grid[m-1][j] == 1 {
			grid[m-1][j] = 0
			dfs(m-1, j)
		}
	}

	var c int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				c++
			}
		}
	}

	fmt.Println(grid)
	return c
}

func main() {
	numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}})
}
