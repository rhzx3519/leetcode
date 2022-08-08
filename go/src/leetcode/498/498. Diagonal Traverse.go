/**
@author ZhengHao Lou
Date    2022/06/14
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/diagonal-traverse/
*/
func findDiagonalOrder(mat [][]int) []int {
	var ans []int
	m, n := len(mat), len(mat[0])

	var f = true
	for i := 0; i < m; i++ {
		var tmp []int
		for k, j := i, 0; k >= 0 && j < n; k, j = k-1, j+1 {
			tmp = append(tmp, mat[k][j])
			//fmt.Println(mat[k][j])
		}
		if !f {
			reverse(tmp)
		}
		ans = append(ans, tmp...)
		f = !f
	}

	for j := 1; j < n; j++ {
		var tmp []int
		for k, i := j, m-1; k < n && i >= 0; k, i = k+1, i-1 {
			tmp = append(tmp, mat[i][k])
			fmt.Println(mat[i][k])
		}
		if !f {
			reverse(tmp)
		}
		ans = append(ans, tmp...)
		f = !f
	}

	fmt.Println(ans)
	return ans
}

func reverse(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

func main() {
	findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
