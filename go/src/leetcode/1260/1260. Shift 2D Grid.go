/**
@author ZhengHao Lou
Date    2022/07/20
*/
package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	for ; k != 0; k-- {
		var tmp = make([]int, m)
		for i := 0; i < m; i++ {
			tmp[i] = grid[i][n-1]
		}

		for j := n - 2; j >= 0; j-- {
			for i := 0; i < m; i++ {
				grid[i][j+1] = grid[i][j]
			}
		}
		for i := 0; i < m; i++ {
			grid[i][0] = tmp[i]
		}

		x := grid[m-1][0]
		for i := m - 2; i >= 0; i-- {
			grid[i+1][0] = grid[i][0]
		}
		grid[0][0] = x
	}
	fmt.Println(grid)
	return grid
}

func main() {
	shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1)
}
