package main

import (
	"math"
)

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxVal := 0
	for l < r {
		maxVal = max(maxVal, (r-l)*min(height[l], height[r]))
		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}
	return maxVal
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
