package main

import (
	"fmt"
	"strings"
)

/*
*
https://leetcode.cn/problems/binary-string-with-substrings-representing-1-to-n/description/
*/
func queryString(s string, n int) bool {
	for i := n>>1 + 1; i <= n; i++ {
		sub := fmt.Sprintf("%b", i)
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(queryString("1", 1))
}
