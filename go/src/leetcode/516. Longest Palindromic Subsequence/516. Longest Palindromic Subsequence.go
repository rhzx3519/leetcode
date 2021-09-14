package main

import "fmt"

/**
https://leetcode-cn.com/problems/longest-palindromic-subsequence/
思路：区间dp
dp[i][j]表示区间[i,j]的最长回文子序列长度, i < j
dp[i][j] = dp[i+1][j-1] + 2, if s[i] == s[j]
dp[i][j] = max(dp[i][j], max(dp[i+1][j], dp[i][j-1])), if s[i] != s[j]
 */
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n-1; i++ {
			j := i + l - 1
			if j >= n {
				break
			}
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j], max(dp[i+1][j], dp[i][j-1]))
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
}
