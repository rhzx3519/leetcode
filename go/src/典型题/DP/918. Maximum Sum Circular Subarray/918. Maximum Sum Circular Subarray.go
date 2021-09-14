package main

import (
	"fmt"
	"math"
)

/**
思路：
环形数组最大值等于
最大子数组和 或 sum - 最小子数组和
 */
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	var sum int
	for _, num := range nums {
		sum += num
	}

	var s int
	var maxVal, minVal = math.MinInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		s = max(0, s) + nums[i]
		maxVal = max(maxVal, s)
	}
	if maxVal < 0 { // nums全是负数
		return maxVal
	}
	s = 0
	for i := 0; i < n; i++ {
		s = min(0, s) + nums[i]
		minVal = min(minVal, s)
	}
	return max(maxVal, sum - minVal)
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

func main() {
	fmt.Println(maxSubarraySumCircular([]int{5,-3,5}))
}