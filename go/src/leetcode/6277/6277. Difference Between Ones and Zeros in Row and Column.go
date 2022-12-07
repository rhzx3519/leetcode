/*
*
@author ZhengHao Lou
Date    2022/11/27
*/
package main

/*
*
https://leetcode.cn/problems/difference-between-ones-and-zeros-in-row-and-column/
*/
func onesMinusZeros(grid [][]int) [][]int {
	var row, col [][]int
	m, n := len(grid), len(grid[0])
	for i := range grid {
		var c0, c1 int
		for j := range grid[i] {
			if grid[i][j] == 0 {
				c0++
			} else {
				c1++
			}
		}
		row = append(row, []int{c0, c1})
	}

	for j := 0; j < n; j++ {
		var c0, c1 int
		for i := 0; i < m; i++ {
			if grid[i][j] == 0 {
				c0++
			} else {
				c1++
			}
		}
		col = append(col, []int{c0, c1})
	}
	diff := make([][]int, m)
	for i := range diff {
		diff[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff[i][j] = row[i][1] + col[j][1] - row[i][0] - col[j][0]
		}
	}
	return diff
}
