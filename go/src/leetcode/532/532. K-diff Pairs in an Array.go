/**
@author ZhengHao Lou
Date    2022/06/16
*/
package main

/**
https://leetcode.cn/problems/k-diff-pairs-in-an-array/
*/
func findPairs(nums []int, k int) int {
	var c int
	counter := make(map[int]int)
	for i := range nums {
		if counter[nums[i]] != 0 && k != 0 {
			continue
		}
		c += counter[nums[i]+k]
		c += counter[nums[i]-k]
		counter[nums[i]]++
	}
	if k == 0 {
		c = 0
		for k := range counter {
			if counter[k] > 1 {
				c++
			}
		}
	}
	return c
}
