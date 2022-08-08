/**
@author ZhengHao Lou
Date    2022/06/27
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/longest-uncommon-subsequence-ii/
*/
func findLUSlength(strs []string) int {
	n := len(strs)
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})
NEXT:
	for i := range strs {
		for j := 0; j < n && len(strs[j]) >= len(strs[i]); j++ {
			if i != j && issub(strs[i], strs[j]) {
				continue NEXT
			}
		}
		return len(strs[i])
	}
	return -1
}

func issub(a, b string) bool {
	na, nb := len(a), len(b)
	var i, j int
	for i < na && j < nb {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	return i == na
}

func main() {
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}))
	fmt.Println(findLUSlength([]string{"aaa", "aaa", "aa"}))
}
