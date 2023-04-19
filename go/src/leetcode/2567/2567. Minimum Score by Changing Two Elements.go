package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/minimum-score-by-changing-two-elements/
1,4,5,7,8
4,4,5,7,7

31,25,72,79,74,65
25,31,65,72,74,79
65,31,65,72,74,65
*/
func minimizeSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	var ans = nums[n-1] - nums[2]
	ans = min(ans, nums[n-3]-nums[0])
	ans = min(ans, nums[n-2]-nums[1])
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimizeSum([]int{1, 4, 3}))
	fmt.Println(minimizeSum([]int{1, 4, 7, 8, 5}))
	fmt.Println(minimizeSum([]int{31, 25, 72, 79, 74, 65}))
}
