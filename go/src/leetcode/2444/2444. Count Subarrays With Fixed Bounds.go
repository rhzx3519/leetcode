/**
@author ZhengHao Lou
Date    2022/10/18
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-subarrays-with-fixed-bounds/
*/
func countSubarrays(nums []int, minK int, maxK int) (tot int64) {
	var l, r1, r2 = -1, -1, -1
	for i := range nums {
		if nums[i] > maxK || nums[i] < minK {
			l = i
		}
		if nums[i] == minK {
			r1 = i
		}
		if nums[i] == maxK {
			r2 = i
		}

		tot += int64(max(0, min(r1, r2)-l))
	}

	return
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

func main() {
	//fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	//fmt.Println(countSubarrays([]int{1, 1, 1, 1}, 1, 1))
	fmt.Println(countSubarrays([]int{35054, 398719, 945315, 945315, 820417, 945315, 35054, 945315,
		171832, 945315, 35054, 109750, 790964, 441974, 552913}, 35054, 945315))
}
