/**
@author ZhengHao Lou
Date    2022/09/26
*/
package main

/**
https://leetcode.cn/problems/longest-subarray-with-maximum-bitwise-and/
*/
func longestSubarray(nums []int) (ans int) {
	var mx, c int
	for i := range nums {
		if nums[i] > mx {
			mx = nums[i]
			c = 1
			ans = 0
		} else if nums[i] == mx {
			c++
		} else {
			c = 0
		}
		ans = max(ans, c)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
