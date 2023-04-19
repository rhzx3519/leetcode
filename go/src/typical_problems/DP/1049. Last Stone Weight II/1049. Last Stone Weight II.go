package main

import "fmt"

/**
https://leetcode-cn.com/problems/last-stone-weight-ii/
思路：题目等价于将石头分成A、B两堆，使得它们之间的差值最小，可以转化为01背包问题，背包容量为S/2, S为所有石头重量之和，
求能装下的最大石头重量。
 */
func lastStoneWeightII(stones []int) int {
	var s int
	for _, stone := range stones {
		s += stone
	}
	var target = s/2
	dp := make([]int, target+1)

	for _, stone := range stones {
		for i := target; i >= stone; i-- {
			dp[i] = max(dp[i], dp[i-stone] + stone)
		}
	}

	return s - dp[target]<<1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2,7,4,1,8,1}))
	fmt.Println(lastStoneWeightII([]int{31,26,33,21,40}))
	fmt.Println(lastStoneWeightII([]int{1,2}))
}


