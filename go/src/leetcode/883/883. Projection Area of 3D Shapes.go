/**
@author ZhengHao Lou
Date    2022/04/26
*/
package main

/**
https://leetcode-cn.com/problems/projection-area-of-3d-shapes/
*/
func projectionArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var s int
	for i := range grid {
		var mx int
		for j := range grid[i] {
			if mx < grid[i][j] {
				mx = grid[i][j]
			}
			if grid[i][j] != 0 {
				s++
			}
		}
		s += mx
	}
	for j := 0; j < n; j++ {
		var mx int
		for i := 0; i < m; i++ {
			if mx < grid[i][j] {
				mx = grid[i][j]
			}
		}
		s += mx
	}
	
	return s
}
