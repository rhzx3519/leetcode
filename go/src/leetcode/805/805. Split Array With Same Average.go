/**
@author ZhengHao Lou
Date    2022/11/14
*/
package main

/**
https://leetcode.cn/problems/split-array-with-same-average/
*/
func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}

	sum := 0
	for _, x := range nums {
		sum += x
	}
	for i := range nums {
		nums[i] = nums[i]*n - sum
	}

	m := n / 2
	left := map[int]bool{}
	for i := 1; i < 1<<m; i++ {
		tot := 0
		for j, x := range nums[:m] {
			if i>>j&1 > 0 {
				tot += x
			}
		}
		if tot == 0 {
			return true
		}
		left[tot] = true
	}

	rsum := 0
	for _, x := range nums[m:] {
		rsum += x
	}
	for i := 1; i < 1<<(n-m); i++ {
		tot := 0
		for j, x := range nums[m:] {
			if i>>j&1 > 0 {
				tot += x
			}
		}
		if tot == 0 || rsum != tot && left[-tot] {
			return true
		}
	}
	return false
}
