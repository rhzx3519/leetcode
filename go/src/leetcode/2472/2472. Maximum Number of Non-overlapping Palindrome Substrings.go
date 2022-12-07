/**
@author ZhengHao Lou
Date    2022/11/16
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/maximum-number-of-non-overlapping-palindrome-substrings/
*/
func maxPalindromes(s string, k int) int {
	n := len(s)
	f := make([]int, n+1)
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		f[l+1] = max(f[l+1], f[l])
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 >= k {
				f[r+1] = max(f[r+1], f[l]+1)
				break
			}
			l--
			r++
		}
	}
	return f[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxPalindromes("abaccdbbd", 3))
}
