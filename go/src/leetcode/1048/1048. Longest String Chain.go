package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/longest-string-chain/
*/
func longestStrChain(words []string) (tot int) {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	n := len(words)
	f := make([]int, n)
	mem := make(map[string]int)
	for i, word := range words {
		f[i] = 1
		for j := 0; j < len(word); j++ {
			if k, ok := mem[word[:j]+word[j+1:]]; ok {
				f[i] = max(f[i], f[k]+1)
			}
		}
		mem[word] = i
		tot = max(tot, f[i])
	}
	fmt.Println(tot, f)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"})
	longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"})
	longestStrChain([]string{"abcd", "dbqca"})
}
