/**
@author ZhengHao Lou
Date    2022/03/14
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-artifacts-that-can-be-extracted/
*/
func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for i := range dig {
		r, c := dig[i][0], dig[i][1]
		grid[r][c] = 1
	}

	var c int
	for i := range artifacts {
		var f = true
		r1, c1, r2, c2 := artifacts[i][0], artifacts[i][1], artifacts[i][2], artifacts[i][3]
		for i := r1; i <= r2; i++ {
			for j := c1; j <= c2; j++ {
				if grid[i][j] == 0 {
					f = false
					break
				}
			}
		}
		if f {
			c++
		}
	}

	fmt.Println(c)
	return c
}

func main() {
	digArtifacts(2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}})
}
