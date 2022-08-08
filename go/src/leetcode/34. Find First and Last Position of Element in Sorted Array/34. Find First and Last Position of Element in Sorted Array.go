package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
 */
func searchRange(nums []int, target int) []int {
	n := len(nums)
	l, r := 0, n
	var m int
	for l < r {
		m = l + (r - l) >> 1
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l == n || nums[l] != target {
		return []int{-1, -1}
	}
	for r = l + 1; r < n && nums[r] == nums[l]; r++ {

	}
	return []int{l, r - 1}
}

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
}