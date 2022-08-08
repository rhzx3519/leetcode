/**
@author ZhengHao Lou
Date    2022/02/01
*/
package main

/**
https://leetcode-cn.com/problems/longest-nice-substring/
*/
func longestNiceSubstring(s string) string {
	n := len(s)
	var l, idx = 0, -1
	for i := range s {
		var a, b int
		for j := i; j < n; j++ {
			if s[j] >= 'a' && s[j] <= 'z' {
				a |= 1 << int(s[j]-'a')
			} else {
				b |= 1 << int(s[j]-'A')
			}
			if a == b && j-i+1 > l {
				l = j - i + 1
				idx = i
			}
		}
	}
	if idx == -1 {
		return ""
	}
	return s[idx : idx+l]
}

func main() {
	longestNiceSubstring("YazaAay")
}
