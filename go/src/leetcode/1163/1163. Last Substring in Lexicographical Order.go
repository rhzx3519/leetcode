package main

/*
*
https://leetcode.cn/problems/last-substring-in-lexicographical-order/description/
*/
func lastSubstring(s string) string {
	n := len(s)
	l, r, step := 0, 1, 0
	for r+step < n {
		if s[l+step] == s[r+step] {
			step++
		} else {
			if s[l+step] < s[r+step] {
				l += step + 1
			} else {
				r += step + 1
			}
			step = 0
			r = max(l+1, r)
		}
	}
	return s[l:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	lastSubstring("leetcode")
}
