/**
@author ZhengHao Lou
Date    2023/01/26
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/minimum-cost-to-split-an-array/description/
思路：动态规划
f[i+1] 表示以nums[:i+1]的最小代价之和
f[i+1] = min(f[j], i-j+1-unique+k), 0<=j<=i, unique表示nums[j:i+1]中只出现一次的数字的数量
*/
func minCost(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		state, unique, v := make([]int, n), 0, math.MaxInt32>>1
		for j := i; j >= 0; j-- {
			x := nums[j]
			if state[x] == 0 {
				state[x] = 1
				unique++
			} else if state[x] == 1 {
				state[x] = 2
				unique--
			}
			v = min(v, f[j]+i-j+1-unique)
		}
		f[i+1] = v + k
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCost([]int{1, 2, 1, 2, 1, 3, 3}, 2))
	fmt.Println(minCost([]int{1, 2, 1, 2, 1}, 2))
	fmt.Println(minCost([]int{1, 2, 1, 2, 1}, 5))
}
