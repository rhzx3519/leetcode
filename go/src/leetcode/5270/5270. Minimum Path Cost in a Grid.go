/**
@author ZhengHao Lou
Date    2022/06/13
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/minimum-path-cost-in-a-grid/
*/
func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			if i == 0 {
				f[i][j] = grid[i][j]
			} else {
				f[i][j] = math.MaxInt32
			}
		}
	}
	
	for k := 0; k < m-1; k++ {
		for l := range grid[k] {
			i := grid[k][l]
			for j := 0; j < n; j++ {
				f[k+1][j] = min(f[k+1][j], grid[k+1][j]+moveCost[i][j]+f[k][l])
			}
		}
		fmt.Println(f[k+1])
	}
	var ans = math.MaxInt32
	for j := 0; j < n; j++ {
		ans = min(ans, f[m-1][j])
	}
	fmt.Println(f)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minPathCost([][]int{{5, 3}, {4, 0}, {2, 1}},
		[][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}})
}
