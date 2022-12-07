/*
*
@author ZhengHao Lou
Date    2022/11/21
*/
package main

/*
*
https://leetcode.cn/problems/number-of-unequal-triplets-in-array/
*/
func unequalTriplets(nums []int) (tot int) {
	n := len(nums)
	for i := range nums {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] != nums[j] && nums[j] != nums[k] && nums[i] != nums[k] {
					tot++
				}
			}
		}
	}
	return tot
}
