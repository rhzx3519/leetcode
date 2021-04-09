package main

import (
	"fmt"
	"math"
)

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = -1
	for i := 1; i < n+1; i++ {
		dp[i] = n
	}
	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			if isPalindrome(s[j:i]) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n]
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func main() {
	fmt.Println(minCut("a"))
	fmt.Println(minCut("aab"))
	fmt.Println(minCut("ab"))
}