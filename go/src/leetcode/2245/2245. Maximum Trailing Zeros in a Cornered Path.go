/**
@author ZhengHao Lou
Date    2022/04/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-trailing-zeros-in-a-cornered-path/
*/
func maxTrailingZeros(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var g2, g5 = make([][]int, m), make([][]int, m)
	for i := range grid {
		g2[i] = make([]int, n)
		g5[i] = make([]int, n)
		for j := range grid[i] {
			g2[i][j] = count(grid[i][j], 2)
			g5[i][j] = count(grid[i][j], 5)
		}
	}

	row2 := make([][]int, m)
	row5 := make([][]int, m)
	for i := 0; i < m; i++ {
		row2[i] = make([]int, n+1)
		row5[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			row2[i][j+1] = row2[i][j] + g2[i][j]
			row5[i][j+1] = row5[i][j] + g5[i][j]
		}
	}
	col2 := make([][]int, n)
	col5 := make([][]int, n)
	for j := 0; j < n; j++ {
		col2[j] = make([]int, m+1)
		col5[j] = make([]int, m+1)
		for i := 0; i < m; i++ {
			col2[j][i+1] = col2[j][i] + g2[i][j]
			col5[j][i+1] = col5[j][i] + g5[i][j]
		}
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tl := col2[j][i+1] + row2[i][j]
			tr := col2[j][i+1] + row2[i][n] - row2[i][j+1]
			bl := col2[j][m] - col2[j][i] + row2[i][j]
			br := col2[j][m] - col2[j][i] + row2[i][n] - row2[i][j+1]

			tl5 := col5[j][i+1] + row5[i][j]
			tr5 := col5[j][i+1] + row5[i][n] - row5[i][j+1]
			bl5 := col5[j][m] - col5[j][i] + row5[i][j]
			br5 := col5[j][m] - col5[j][i] + row5[i][n] - row5[i][j+1]

			tl = min(tl, tl5)
			tr = min(tr, tr5)
			bl = min(bl, bl5)
			br = min(br, br5)
			ans = max(ans, max(max(tl, tr), max(bl, br)))
		}
	}

	fmt.Println(grid)
	fmt.Println(row2)
	fmt.Println(col2)
	fmt.Println(ans)
	return ans
}

func count(x, base int) int {
	var c int
	for x%base == 0 {
		x /= base
		c++
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxTrailingZeros([][]int{{23, 17, 15, 3, 20}, {8, 1, 20, 27, 11}, {9, 4, 6, 2, 21}, {40, 9, 1, 10, 6}, {22, 7, 4, 5, 3}})
	maxTrailingZeros([][]int{{899, 727, 165, 249, 531, 300, 542, 890}, {981, 587, 565, 943, 875, 498, 582, 672}, {106, 902, 524, 725, 699, 778, 365, 220}})
}
