/**
@author ZhengHao Lou
Date    2022/03/09
*/
package main

/**
https://leetcode-cn.com/problems/smallest-rotation-with-highest-score/
*/
func bestRotation(nums []int) int {
	n := len(nums)
	diff := make([]int, n)
	for i := range nums {
		low := (i + 1) % n
		high := (i - nums[i] + n + 1) % n
		diff[low]++
		diff[high]--
		if low >= high {
			diff[0]++
		}
	}

	var score, mxScore, idx int
	for i := range diff {
		score += diff[i]
		if score > mxScore {
			idx = i
			mxScore = score
		}
	}

	return idx
}
