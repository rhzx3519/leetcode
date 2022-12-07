/**
@author ZhengHao Lou
Date    2022/08/15
*/
package main

/**
https://leetcode.cn/problems/largest-local-values-in-a-matrix/
*/
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	matrix := make([][]int, n-2)
	for i := range matrix {
		matrix[i] = make([]int, n-2)
	}
	// i = l+1, l = i-1, i=1
	for i := 1; i <= n-2; i++ {
		for j := 1; j <= n-2; j++ {
			var mx int
			for k := i - 1; k <= i+1; k++ {
				for z := j - 1; z <= j+1; z++ {
					if grid[k][z] > mx {
						mx = grid[k][z]
					}
				}
			}
			matrix[i-1][j-1] = mx
		}
	}
	return matrix
}

func main() {
	largestLocal([][]int{})
}
