/**
@author ZhengHao Lou
Date    2022/06/06
*/
package main

/**
https://leetcode.cn/problems/min-max-game/
*/
func minMaxGame(nums []int) int {
	for n := len(nums); n != 1; n = len(nums) {
		var tmp []int
		for i := 0; i < n>>1; i++ {
			if i&1 == 0 {
				tmp = append(tmp, min(nums[i<<1], nums[i<<1+1]))
			} else {
				tmp = append(tmp, max(nums[i<<1], nums[i<<1+1]))
			}
		}
		nums = tmp
	}
	return nums[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
