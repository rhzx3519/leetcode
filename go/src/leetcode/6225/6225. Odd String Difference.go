/**
@author ZhengHao Lou
Date    2022/10/30
*/
package main

import (
	"fmt"
)

/**
https://leetcode.cn/problems/odd-string-difference/
*/
func oddString(words []string) string {
	n := len(words[0])
	mp := make(map[string][]int)
	for j, w := range words {
		var d string
		for i := 1; i < n; i++ {
			d += fmt.Sprintf(",%v", int(w[i]-w[i-1]))
		}
		mp[d] = append(mp[d], j)
	}
	for i := range mp {
		if len(mp[i]) == 1 {
			return words[mp[i][0]]
		}
	}
	return ""
}

func main() {
	oddString([]string{"abm", "bcn", "alm"})
}
