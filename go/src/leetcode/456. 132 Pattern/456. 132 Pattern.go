package main

import "math"

func find132pattern(nums []int) bool {
	minVal := math.MinInt32
	st := []int{}
	for i := len(nums)-1; i >= 0; i-- {
		if minVal > nums[i] {
			return true
		}
		for len(st)>0 && st[len(st)-1]<nums[i] {
			minVal = st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, nums[i])
	}
	return false
}
