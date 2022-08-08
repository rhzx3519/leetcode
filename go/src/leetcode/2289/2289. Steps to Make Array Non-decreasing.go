/**
@author ZhengHao Lou
Date    2022/05/30
*/
package main

import (
	"fmt"
)

/**
https://leetcode.cn/problems/steps-to-make-array-non-decreasing/
*/
func totalSteps(nums []int) int {
	n := len(nums)
	f := make([]int, n)

	var ans int
	var st []int
	for i := 0; i < n; i++ {
		var p int
		for len(st) != 0 && nums[st[len(st)-1]] <= nums[i] {
			p = max(p, f[st[len(st)-1]])
			st = st[:len(st)-1]
		}
		if len(st) != 0 {
			f[i] = p + 1
			ans = max(ans, f[i])
		}
		st = append(st, i)
	}

	fmt.Println(f, ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(totalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}))
	//fmt.Println(totalSteps([]int{4, 5, 7, 7, 13}))
	//fmt.Println(totalSteps([]int{5, 11, 7, 8, 11}))
	fmt.Println(totalSteps([]int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3})) // 6
}
