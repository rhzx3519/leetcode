package main

import "fmt"

/*
*
https://leetcode.cn/problems/decrease-elements-to-make-array-zigzag/
*/
func movesToMakeZigzag(nums []int) int {
	n := len(nums)
	var odd, even int
	for i := range nums {
		if i&1 == 1 { // i-1, i+1 is even
			var l, r int
			if i-1 >= 0 && nums[i] >= nums[i-1] {
				l = nums[i] - nums[i-1] + 1
			}
			if i+1 < n && nums[i] >= nums[i+1] {
				r = nums[i] - nums[i+1] + 1
			}
			odd += max(l, r)
		} else { // i-1, i+1 is odd
			var l, r int
			if i-1 >= 0 && nums[i] >= nums[i-1] {
				l = nums[i] - nums[i-1] + 1
			}
			if i+1 < n && nums[i] >= nums[i+1] {
				r = nums[i] - nums[i+1] + 1
			}
			even += max(l, r)
		}
	}
	return min(odd, even)
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
	fmt.Println(movesToMakeZigzag([]int{9, 6, 1, 6, 2}))
}
