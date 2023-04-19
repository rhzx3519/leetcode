package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced/
baab
bbab
*/
func minimumDeletions(s string) int {
	n := len(s)
	left, right := make([]int, n), make([]int, n)
	var cb, ca int
	for i := range s {
		left[i] = cb
		if s[i] == 'b' {
			cb++
		}
	}
	for i := n - 1; i >= 0; i-- {
		right[i] = ca
		if s[i] == 'a' {
			ca++
		}
	}

	var ans = n
	for i := 0; i < n; i++ {
		ans = min(ans, left[i]+right[i])
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumDeletions("aababbab"))
	fmt.Println(minimumDeletions("bbaaaaabb"))
}
