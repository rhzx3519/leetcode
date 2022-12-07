/**
@author ZhengHao Lou
Date    2022/11/07
*/
package main

/**
https://leetcode.cn/problems/apply-operations-to-an-array/
*/
func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] <<= 1
			nums[i+1] = 0
		}
	}
	n := len(nums)
	var ans = make([]int, n)
	var j int
	for i := range nums {
		if nums[i] != 0 {
			ans[j] = nums[i]
			j++
		}
	}
	return ans
}
