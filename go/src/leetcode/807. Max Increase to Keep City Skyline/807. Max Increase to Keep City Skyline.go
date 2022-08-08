/**
@author ZhengHao Lou
Date    2021/12/13
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/
 */
func maxIncreaseKeepingSkyline(grid [][]int) int {
	var ans int
	m, n := len(grid), len(grid[0])
	row, col := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		var x int
		for j := 0; j < n; j++ {
			x = max(x, grid[i][j])
		}
		row[i] = x
	}
	for j := 0; j < n; j++ {
		var x int
		for i := 0; i < n; i++ {
			x = max(x, grid[i][j])
		}
		col[j] = x
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans += min(row[i], col[j])	- grid[i][j]
		}
	}

	fmt.Println(ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	maxIncreaseKeepingSkyline([][]int{{3,0,8,4},{2,4,5,7},{9,2,6,3},{0,3,1,0}})
}
