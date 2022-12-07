/**
@author ZhengHao Lou
Date    2022/10/05
*/
package main

/**
https://leetcode.cn/problems/maximum-sum-of-an-hourglass/
*/
func maxSum(grid [][]int) (mx int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i+2 < m; i++ {
		for j := 0; j+2 < n; j++ {
			var s int
			for k := 0; k < 3; k++ {
				s += grid[i][j+k]
			}
			s += grid[i+1][j+1]
			for k := 0; k < 3; k++ {
				s += grid[i+2][j+k]
			}
			if s > mx {
				mx = s
			}
		}
	}
	return
}
