/**
@author ZhengHao Lou
Date    2022/07/05
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/number-of-increasing-paths-in-a-grid/
*/
const MOD int = 1e9 + 7

var (
	dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
)

func countPaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if mem[x][y] != 0 {
			return mem[x][y]
		}
		var c = 1
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] > grid[x][y] {
				c = (c + dfs(nx, ny)) % MOD
			}
		}
		mem[x][y] = c
		return c
	}

	var c int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c = (c + dfs(i, j)) % MOD
		}
	}

	fmt.Println(c)
	return c
}

func main() {
	countPaths([][]int{{1, 1}, {3, 4}})
	countPaths([][]int{{1}, {2}})
}
