package main

import (
	"fmt"
)

/*
*
https://leetcode.cn/problems/matrix-similarity-after-cyclic-shifts/
*/
func areSimilar(mat [][]int, k int) bool {
	_, n := len(mat), len(mat[0])
	k = k % n
	if k == 0 {
		return true
	}

	for i := range mat {
		if i&1 == 0 {
			for j := range mat[i] {
				if j < n-k {
					if mat[i][j] != mat[i][j+k] {
						return false
					}
				} else {
					if mat[i][j] != mat[i][j-(n-k)] {
						return false
					}
				}
			}
		} else {
			for j := range mat[i] {
				if j < k {
					if mat[i][j] != mat[i][n-k+j] {
						return false
					}
				} else {
					if mat[i][j] != mat[i][j-k] {
						return false
					}
				}
			}
		}
	}

	return true
}

func main() {
	//fmt.Println(areSimilar([][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}, 2))
	//fmt.Println(areSimilar([][]int{{7, 7}, {10, 10}, {4, 4}}, 2))
	fmt.Println(areSimilar([][]int{{8, 8}, {6, 6}, {2, 2}, {8, 8}, {9, 9}, {8, 8},
		{10, 10}, {3, 3}, {4, 4}, {5, 5}}, 1))
}
