package main

import "fmt"

//# 思路：dp[i][j]表示s[i:j+1]是一个回文串, dp[i][j] = (dp[i+1][j-1] or j-i<3) and s[i]==s[j]
//# 遍历方向是斜上角方向，所以i是从n-1到0，j是从i+1到n-1
//# time: O(N^2), space: O(N^2)

func longestPalindrome(s string) string {
	var maxLen int
	var ans string

	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for l := 1; l <= n; l++ {
		for i := 0; i + l <= n; i++ {
			// l = j - i + 1
			j := i + l - 1
			if s[i] == s[j] {
				if l < 3 || dp[i+1][j-1] {
					dp[i][j] = true
					if l > maxLen {
						maxLen = l
						ans = s[i:j+1]
					}
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(longestPalindrome("babad"))	// bab
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
}