package main

// https://leetcode-cn.com/problems/minimum-limit-of-balls-in-a-bag/solution/dai-zi-li-zui-shao-shu-mu-de-qiu-by-zero-upwe/

import "math"

func minimumSize(nums []int, maxOperations int) int {
	l, r := 1, max(nums)
	ans := r
	for l <= r {
		y := l + (r - l)/2
		ops := 0
		for _, x := range nums {
			ops += (x - 1) / y
		}
		if ops <= maxOperations {
			ans = y
			r = y - 1
		} else {
			l = y + 1
		}
	}
	return ans
}

func max(nums []int) int {
	maxVal := math.MinInt32
	for _, num := range nums {
		if maxVal < num {
			maxVal = num
		}
	}
	return maxVal
}

