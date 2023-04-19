package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		c0, c1 := count(s)
		for i := m; i >= c0; i-- {
			for j := n; j >= c1; j-- {
				dp[i][j] = max(dp[i][j], dp[i-c0][j-c1]+1)
			}
		}
	}
	return dp[m][n]
}

func count(s string) (c0, c1 int) {
	for i := range s {
		if s[i] == '0' {
			c0++
		} else {
			c1++
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}