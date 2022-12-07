/**
@author ZhengHao Lou
Date    2022/10/25
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/shortest-bridge/
*/
var diff = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var que [][]int
	var vis = make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		for _, d := range diff {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] && grid[nx][ny] == 1 {
				que = append(que, []int{nx, ny})
				vis[nx][ny] = true
				grid[nx][ny] = 2
				dfs(nx, ny)
			}
		}
	}

OUT:
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				vis[i][j] = true
				grid[i][j] = 2
				que = append(que, []int{i, j})
				dfs(i, j)
				break OUT
			}
		}
	}

	var bfs func() int
	bfs = func() int {
		for step := 0; len(que) != 0; step++ {
			for tmp := len(que); tmp != 0; tmp-- {
				x, y := que[0][0], que[0][1]
				que = que[1:]
				if grid[x][y] == 1 {
					return step - 1
				}
				for _, d := range diff {
					nx, ny := x+d[0], y+d[1]
					if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] {
						vis[nx][ny] = true
						que = append(que, []int{nx, ny})
					}
				}
			}
		}
		return -1
	}

	return bfs()
}

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1}, {1, 0}}))
	fmt.Println(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
}
