package main

/*
*	https://leetcode.cn/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip/
 */
func isPossibleToCutPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == m-1 && j == n-1 {
			return true
		}
		grid[i][j] = 0
		return i+1 < m && grid[i+1][j] == 1 && dfs(i+1, j) ||
			j+1 < n && grid[i][j+1] == 1 && dfs(i, j+1)
	}

	return !dfs(0, 0) || !dfs(0, 0)
}
