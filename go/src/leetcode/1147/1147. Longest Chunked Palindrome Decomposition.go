package main

/*
*
https://leetcode.cn/problems/longest-chunked-palindrome-decomposition/description/
*/
func longestDecomposition(s string) (tot int) {
	for s != "" {
		i, n := 1, len(s)
		for i <= n>>1 && s[:i] != s[n-i:] {
			i++
		}
		if i > n/2 { // 无法分割
			tot++
			break
		}
		tot += 2
		s = s[i : n-i] // 分割出s[:i], s[n-i:]
	}
	return
}
