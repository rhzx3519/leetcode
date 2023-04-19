package main

import "fmt"

/**
回文串编辑
给定一个字符串 str, 每一次操作都可以在任意位置插入一个字符或者删去某个字符
求可以让 str 变为回文串的最少的操作次数.
注: 回文串是正读和反读都相同的字符串; str 的所有字母都是小写.
难度
Hard
Example
input: aabbb; output: 1(str 可变为 aaabbb 或者 aabb)
input: aa; output: 0
input: tiger; output: 4(str 可变为 tigeregit)
 */

func palindromeEdit(s string) int {
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
	fmt.Println(palindromeEdit("aabbb"))
	//fmt.Println(palindromeEdit("aa"))
	//fmt.Println(palindromeEdit("tiger"))
}
