package main

import "fmt"

/**
https://leetcode-cn.com/problems/target-sum/
思路：根据题设可以讲数组分为正子集P和负子集N, P-N=S, P+N=T, T为数组所有元素之和
可以推出P = (T+S)/2, 题目转化为从数组中选取若干个元素，使得和为(T+S)/2
 */
func findTargetSumWays(nums []int, target int) int {
	var s int
	for _, num := range nums {
		s += num
	}
	if s < target || (s + target) & 1 == 1 {
		return 0
	}
	var t = (s + target)/2
	dp := make([]int, t+1)
	dp[0] = 1
	for _, num := range nums {
		for i := t; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[t]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1,1,1,1,1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))
}
