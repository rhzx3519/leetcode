package main

import "fmt"

/**
https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
思路：
dp[i][j] 表示[i,j]变成回文串的最少插入次数，
dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1, if s[i] != s[j]
dp[i][j] = min(dp[i+1][j] + 1, dp[i][j-1] + 1, dp[i+1][j-1]), if s[i] != s[j]

 */
func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i >= j {
				dp[i][j] = 0
			} else {
				dp[i][j] = n
			}
		}
	}

	for span := 2; span <= n; span++ {
		for i := 0; i <= n - span; i++ {
			j := i + span - 1
			dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			if s[i] == s[j] {
				dp[i][j] = min(dp[i][j], dp[i+1][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minInsertions("zzazz"))
	fmt.Println(minInsertions("mbadm"))
	fmt.Println(minInsertions("leetcode"))
	fmt.Println(minInsertions("g"))
}
