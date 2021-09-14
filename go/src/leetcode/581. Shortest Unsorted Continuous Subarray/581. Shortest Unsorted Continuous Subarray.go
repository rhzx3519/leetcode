package main

import (
	"fmt"
	"math"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	maxVal, minVal := math.MinInt32, math.MaxInt32
	var l, r int
	for i := 0; i < n; i++ {
		maxVal = max(maxVal, nums[i])
		if nums[i] < maxVal {
			r = i
		}
	}
	for i := n-1; i >= 0; i-- {
		minVal = min(minVal, nums[i])
		if nums[i] > minVal {
			l = i
		}
	}
	if l==0 && r==0 {
		return 0
	}
	fmt.Println(r - l + 1)
	return r - l + 1
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
	findUnsortedSubarray([]int{2,6,4,8,10,9,15})
	findUnsortedSubarray([]int{1,2,3,4})
	findUnsortedSubarray([]int{1})
	findUnsortedSubarray([]int{2,3,3,2,4})
	findUnsortedSubarray([]int{1,2,3,3,3})
	findUnsortedSubarray([]int{1,2,4,5,3})
}