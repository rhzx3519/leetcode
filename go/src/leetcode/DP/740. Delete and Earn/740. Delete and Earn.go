package main
/**
思路：动态规划
dp[i]表示删除或者不删除i点数时，最大能获取的点数
 */
func deleteAndEarn(nums []int) int {
	const M = 1001
	dp := make([]int, M)
	a := make([]int, M)
	for _, num := range nums {
		a[num] += num
	}
	dp[0] = a[0]
	dp[1] = max(a[0], a[1])
	for i := 2; i < M; i++ {
		dp[i] = max(dp[i-1], dp[i-2] + a[i])
	}
	return dp[M-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
