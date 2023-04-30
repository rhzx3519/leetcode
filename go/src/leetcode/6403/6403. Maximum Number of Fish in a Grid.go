package main

import "fmt"

/*
*
https://leetcode.cn/problems/maximum-number-of-fish-in-a-grid/
*/
var offset = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(x, y, color int) int
	dfs = func(x, y, color int) int {
		var tot = grid[x][y]
		grid[x][y] = color
		for _, d := range offset {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] > 0 {
				tot += dfs(nx, ny, color)
			}
		}
		return tot
	}

	var ans, color = 0, -1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				tot := dfs(i, j, color)
				color--
				if tot > ans {
					ans = tot
				}
			}
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	//findMaxFish([][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}})
	findMaxFish([][]int{{0, 3, 8}})
}
