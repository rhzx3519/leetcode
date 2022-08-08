/**
@author ZhengHao Lou
Date    2021/12/03
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/
思路：动态规划
f[i+1][j+1]表示以(i,j)为右下角的最大正方形边长
f[i+1][j+1] = min(f[i][j], min(f[i][j+1], f[i+1][j])) + 1, if matrix[i][j] == 1
 */
func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m + 1)
	for i := range f {
		f[i] = make([]int, n + 1)
	}

	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				f[i+1][j+1] = min(f[i][j], min(f[i][j+1], f[i+1][j])) + 1
				count += f[i+1][j+1]
			}
		}
	}

	fmt.Println(f)
	fmt.Println(count)
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	countSquares([][]int{{0,1,1,1},{1,1,1,1},{0,1,1,1}})
}
