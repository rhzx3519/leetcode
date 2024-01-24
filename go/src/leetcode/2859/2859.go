package main

import "math/bits"

/*
*
https://leetcode.cn/problems/sum-of-values-at-indices-with-k-set-bits/?envType=daily-question&envId=2024-01-25
*/
func sumIndicesWithKSetBits(nums []int, k int) (tot int) {
	for i := range nums {
		if bits.OnesCount(uint(i)) == k {
			tot += nums[i]
		}
	}
	return
}
