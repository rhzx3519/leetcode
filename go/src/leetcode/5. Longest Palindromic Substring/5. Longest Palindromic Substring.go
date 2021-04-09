package main

//# 思路：dp[i][j]表示s[i:j+1]是一个回文串, dp[i][j] = (dp[i+1][j-1] or j-i<3) and s[i]==s[j]
//# 遍历方向是斜上角方向，所以i是从n-1到0，j是从i+1到n-1
//# time: O(N^2), space: O(N^2)

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	l, r := 0, 0
	for i := n-1; i >= 0; i-- {
		dp[i][i] = true
		for j := i+1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (dp[i+1][j-1] || j-i < 3)
			if dp[i][j] && j-i > r-l {
				l, r = i, j
			}
		}
	}

	return s[l:r+1]
}
