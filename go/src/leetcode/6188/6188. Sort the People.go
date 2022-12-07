/**
@author ZhengHao Lou
Date    2022/09/26
*/
package main

import "sort"

/**
https://leetcode.cn/problems/sort-the-people/
*/
func sortPeople(names []string, heights []int) []string {
	var nums []int
	for i := range names {
		nums = append(nums, i)
	}
	sort.Slice(nums, func(i, j int) bool {
		return heights[nums[i]] > heights[nums[j]]
	})
	var ans []string
	for _, i := range nums {
		ans = append(ans, names[i])
	}
	return ans
}
