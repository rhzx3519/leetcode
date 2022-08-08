/**
@author ZhengHao Lou
Date    2022/04/02
*/
package main

/**
https://leetcode-cn.com/contest/biweekly-contest-75/problems/find-triangular-sum-of-an-array/
*/
func triangularSum(nums []int) int {
	for len(nums) != 1 {
		var arr = make([]int, len(nums)-1)
		for i := 0; i < len(nums)-1; i++ {
			arr[i] = (nums[i] + nums[i+1]) % 10
		}
		nums = arr
	}
	return nums[0]
}
