package main

import "fmt"

func maxAscendingSum(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	s := nums[0]
	maxVal := s
	for i:= 1; i < n; i++ {
		if i > 0 && nums[i] > nums[i-1] {
			s += nums[i]
		} else {
			s = nums[i]
		}
		maxVal = max(maxVal, s)
	}
	return maxVal
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxAscendingSum([]int{10,20,30,5,10,50}))
	fmt.Println(maxAscendingSum([]int{10,20,30,40,50}))
}
