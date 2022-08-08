/**
@author ZhengHao Lou
Date    2022/01/09
*/
package main

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	for i := 0; i < n; i++ {
		var f1 = make([]int, n+1)
		var f2 = make([]int, n+1)

		for j := 0; j < n; j++ {
			f1[matrix[i][j]] = 1
			f2[matrix[j][i]] = 1
		}
		for k := 1; k <= n; k++ {
			if f1[k] == 0 || f2[k] == 0 {
				return false
			}
		}
	}

	return true
}
