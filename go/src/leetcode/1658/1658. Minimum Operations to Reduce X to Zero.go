/**
@author ZhengHao Lou
Date    2023/01/07
*/
package main

import (
	"fmt"
)

/**
https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/
思路：滑动窗口
*/
func minOperations(nums []int, x int) int {
	n := len(nums)
	var s int
	for i := range nums {
		s += nums[i]
	}
	if s < x {
		return -1
	}
	if s == x {
		return n
	}
	
	var ans = n
	var l, c int
	for r := 0; r < n; r++ {
		c += nums[r]
		for l < r && s-c < x {
			c -= nums[l]
			l++
		}
		if s-c == x {
			ans = min(ans, n-(r-l+1))
		}
	}
	if ans == n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(minOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
}
