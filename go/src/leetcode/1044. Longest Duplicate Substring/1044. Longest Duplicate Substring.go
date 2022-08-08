package main

import "fmt"

/**
https://leetcode-cn.com/problems/longest-duplicate-substring/solution/gong-shui-san-xie-zi-fu-chuan-ha-xi-ying-hae9/
思路：字符串哈希，二分
time: O(nlogn)
space: O(n)
 */
const (
	N int = 3e4 + 10
	P int = 131313
)
var (
	p []int = make([]int, N)
	h []int = make([]int, N)
)
func longestDupSubstring(s string) string {
	n := len(s)
	p[0] = 1
	for i := 1; i <= n; i++ {
		h[i] = h[i-1]*P + int(s[i-1])
		p[i] = p[i-1]*P
	}

	var ans = ""
	l, r := 0, n
	var m int
	for l < r {
		m = l + (r - l)>>1
		sub := check(s, m)
		if len(sub) == 0 {
			r = m
		} else {
			l = m + 1
		}
		if len(sub) > len(ans) {
			ans = sub
		}
	}

	return ans
}

func check(s string, l int) string {
	n := len(s)
	mapper := make(map[int]int)
	for i := 1; i + l - 1 <= n; i++ {
		var j = i + l - 1
		hash := h[j] - h[i-1]*p[j-i+1]
		if mapper[hash] > 0 {
			return s[i-1:j]
		}
		mapper[hash]++
	}
	return ""
}

func main() {
	fmt.Println(longestDupSubstring("banana"))
	fmt.Println(longestDupSubstring("abcd"))
	fmt.Println(longestDupSubstring("aa"))
}
