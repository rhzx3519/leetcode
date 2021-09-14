package main

import (
	"math"
)

func maxProduct(nums []int) int {
	// f1保存当前连续数组的最大值, f2保存当前连续数组的最小值
	f1, f2 := 1, 1
	maxVal := math.MinInt32;
	for _, num := range nums {
		if num < 0 {
			f1, f2 = f2, f1
		}
		f1 = max(num, f1*num)
		f2 = min(num, f2*num)

		maxVal = max(maxVal, f1)
	}

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
