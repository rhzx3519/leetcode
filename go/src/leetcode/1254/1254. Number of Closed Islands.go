package main

/*
*
https://leetcode.cn/problems/number-of-closed-islands/description/
*/

func closedIsland(grid [][]int) (tot int) {
	m, n := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 0 {
			return
		}
		grid[x][y] = 1
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				tot++
				dfs(i, j)
			}
		}
	}
	return tot
}

func main() {
	closedIsland([][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}})
}
