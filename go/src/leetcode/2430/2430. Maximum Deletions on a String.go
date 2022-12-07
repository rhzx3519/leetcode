/**
@author ZhengHao Lou
Date    2022/10/05
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/maximum-deletions-on-a-string/
思路：线性动态规划
*/
func deleteString(s string) int {
	n := len(s)
	if oneLetter(s) {
		return n
	}
	lcp := make([][]int, n+1)
	lcp[n] = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		lcp[i] = make([]int, n+1)
		for j := n - 1; j > i; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}
	f := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := 1; i+j*2 <= n; j++ {
			if lcp[i][i+j] >= j { // s[i:i+j] == s[i+j:i+2*j]
				f[i] = max(f[i], f[i+j])
			}
		}
		f[i]++
	}

	return f[0]
}

func oneLetter(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(deleteString("abcabcdabc"))
	fmt.Println(deleteString("aaabaab"))
	fmt.Println(deleteString("aaaaa"))
}
