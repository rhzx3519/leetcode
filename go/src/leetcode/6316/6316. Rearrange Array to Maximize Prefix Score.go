package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/rearrange-array-to-maximize-prefix-score/
*/
func maxScore(nums []int) (tot int) {
	n := len(nums)
	prefix := make([]int, n+1)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i := range nums {
		prefix[i+1] = nums[i] + prefix[i]
	}
	for i := 1; i <= n; i++ {
		if prefix[i] <= 0 {
			break
		}
		tot++
	}
	fmt.Println(prefix, tot)
	return
}

func main() {
	maxScore([]int{2, -1, 0, 1, -3, 3, -3})
	maxScore([]int{-2, -3})
}
