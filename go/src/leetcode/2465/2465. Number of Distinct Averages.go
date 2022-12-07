/**
@author ZhengHao Lou
Date    2022/11/15
*/
package main

import (
	"sort"
)

/**
https://leetcode.cn/problems/number-of-distinct-averages/
*/
func distinctAverages(nums []int) int {
	var counter = make(map[int]int)
	sort.Ints(nums)
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		counter[nums[l]+nums[r]]++
	}
	return len(counter)
}
