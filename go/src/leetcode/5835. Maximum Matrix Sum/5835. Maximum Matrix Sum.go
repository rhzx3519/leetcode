package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/maximum-matrix-sum/
 */
func maxMatrixSum(matrix [][]int) int64 {
	var s, neg int64
	var minVal = math.MaxInt32
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < 0 {
				neg++
			}
			if (abs(matrix[i][j]) < minVal) {
				minVal = abs(matrix[i][j])
			}
			s += int64(abs(matrix[i][j]))
		}
	}
	if neg&1 == 1 && minVal != math.MaxInt32 {
		s -= int64(minVal)<<1
	}
	return s
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxMatrixSum([][]int{{1,-1},{-1,1}}))
	fmt.Println(maxMatrixSum([][]int{{1,2,3},{-1,-2,-3},{1,2,3}}))
}
