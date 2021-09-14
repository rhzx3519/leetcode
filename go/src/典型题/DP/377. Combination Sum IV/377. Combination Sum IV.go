package main

/**
思路：完全背包问题
dp[i]表示nums中数字和为i的组合个数
time: O(n*target), space: O(target)
 */
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
