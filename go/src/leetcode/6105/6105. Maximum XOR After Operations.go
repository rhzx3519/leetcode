/**
@author ZhengHao Lou
Date    2022/06/27
*/
package main

/**
https://leetcode.cn/problems/maximum-xor-after-operations/
*/
func maximumXOR(nums []int) int {
	var x int
	for i := range nums {
		x |= nums[i]
	}
	return x
}
