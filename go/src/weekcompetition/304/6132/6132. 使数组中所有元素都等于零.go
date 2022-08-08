/**
@author ZhengHao Lou
Date    2022/07/31
*/
package main

func minimumOperations(nums []int) int {
	var step int
	for ; len(nums) != 0; step++ {
		var mi = 101
		for i := range nums {
			if nums[i] > 0 && nums[i] < mi {
				mi = nums[i]
			}
		}
		if mi == 101 {
			return step
		}
		var tmp []int
		for i := range nums {
			if nums[i]-mi > 0 {
				tmp = append(tmp, nums[i]-mi)
			}
		}
		nums = tmp
	}
	return step
}
