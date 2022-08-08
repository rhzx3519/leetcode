/**
@author ZhengHao Lou
Date    2021/10/25
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
 */
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	nums := []int{}
	for i := 0; i < m; i++ {
		nums = append(nums, matrix[i][0])
	}
	r := upperBound(nums, target)
	for i := 0; i < r; i++ {
		c := lowerBound(matrix[i], target)
		if c < n && matrix[i][c] == target {
			return true
		}
	}
	return false
}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r - l) >> 1
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r - l) >> 1
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	//fmt.Println(searchMatrix([][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}, 5))
	fmt.Println(searchMatrix([][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}, 20))
}