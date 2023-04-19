package main

/**
https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
 */
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n - 1
	var m int
	for l <= r {
		m = l + (r - l)>>1
		if nums[m] == target {
			return m
		}
		if nums[m] < nums[r] { 	// [m, r] 递增
			if nums[m] < target && target <= nums[r] {	// nums[m] < target <= nums[r]
				l = m + 1
			} else {
				r = m - 1
			}
		} else {	// [l, m] 递增
			if nums[l] <= target && target < nums[m] {			// nums[l] <= target < nums[m]
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
