/**
@author ZhengHao Lou
Date    2022/06/27
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/check-if-matrix-is-x-matrix/
*/
func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := grid[i][j]
			if i == j || i+j == n-1 { // 3,0 2,1 1,2
				fmt.Println(i, j)
				if x == 0 {
					return false
				}
			} else {
				if x != 0 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	checkXMatrix([][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}})
}
