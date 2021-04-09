package main

import "fmt"

func numDistinct(s string, t string) int {
	ls, lt := len(s), len(t)
	dp := make([][]int, ls+1)
	for i := 0; i < ls+1; i++ {
		dp[i] = make([]int, lt+1)
		dp[i][0] = 1
	}

	for i := 1; i < ls+1; i++ {
		for j := 1; j < lt+1; j++ {
			if i < j {
				continue
			}
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[ls][lt]
}

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}
