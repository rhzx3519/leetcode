/**
@author ZhengHao Lou
Date    2022/07/18
*/
package main

/**
https://leetcode.cn/problems/maximum-number-of-pairs-in-array/
*/
func numberOfPairs(nums []int) []int {
	var counter = make(map[int]int)
	for i := range nums {
		counter[nums[i]]++
	}
	var a, b int
	for _, v := range counter {
		a += v >> 1
		b += v & 1
	}
	return []int{a, b}
}
