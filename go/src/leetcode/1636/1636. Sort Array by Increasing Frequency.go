/**
@author ZhengHao Lou
Date    2022/09/19
*/
package main

import "sort"

/**
https://leetcode.cn/problems/sort-array-by-increasing-frequency/
*/
func frequencySort(nums []int) []int {
	counter := make(map[int]int)
	for i := range nums {
		counter[nums[i]]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if counter[nums[i]] != counter[nums[j]] {
			return counter[nums[i]] < counter[nums[j]]
		}
		return nums[i] > nums[j]
	})
	return nums
}
