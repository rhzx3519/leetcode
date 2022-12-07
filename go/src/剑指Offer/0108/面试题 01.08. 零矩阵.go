/**
@author ZhengHao Lou
Date    2022/09/30
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/zero-matrix-lcci/
*/
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rows, cols := make([]bool, m), make([]bool, n)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	fmt.Println(rows, cols)

	for i := range rows {
		if rows[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := range cols {
		if cols[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Println(matrix)
}

func main() {
	//[[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	setZeroes([][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5}})
}
