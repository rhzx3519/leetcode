package main

import (
	"bytes"
)

/*
*
https://leetcode.cn/problems/find-the-string-with-lcp/
*/
func findTheString(lcp [][]int) string {
	n := len(lcp)
	s := make([]byte, n)
	for c := 'a'; c <= 'z'; c++ {
		i := bytes.IndexByte(s, 0)
		if i == -1 { // 构造完毕
			break
		}
		for j := i; j < n; j++ {
			if lcp[i][j] > 0 {
				s[j] = byte(c)
			}
		}
	}
	if i := bytes.IndexByte(s, 0); i >= 0 {
		return ""
	}

	// 验证字符串是否符合lcp
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			var actualLCP int
			if s[i] == s[j] {
				actualLCP = 1
				if i+1 < n && j+1 < n {
					actualLCP += lcp[i+1][j+1]
				}
			}
			if lcp[i][j] != actualLCP {
				return ""
			}
		}
	}

	return string(s)
}
