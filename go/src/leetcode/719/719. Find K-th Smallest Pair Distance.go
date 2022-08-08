/**
@author ZhengHao Lou
Date    2022/06/15
*/
package main

import "sort"

/**
https://leetcode.cn/problems/find-k-th-smallest-pair-distance/
*/
func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	
	var getLess = func(d int) int {
		var l, t int
		for r := 0; r < n; r++ {
			for nums[r]-nums[l] > d {
				l++
			}
			t += r - l
		}
		return t
	}
	
	l, r := 0, nums[n-1]-nums[0]+1
	var m int
	for l < r {
		m = l + (r-l)>>1
		if getLess(m) >= k {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
