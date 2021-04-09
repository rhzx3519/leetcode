package main

func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 0; i < num+1; i++ {
		dp[i] += dp[i>>1] + i&1
	}
	return dp
}
