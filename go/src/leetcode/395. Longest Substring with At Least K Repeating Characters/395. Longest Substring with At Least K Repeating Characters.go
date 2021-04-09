package main

import (
	"fmt"
	"math"
	"strings"
)

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	st := make(map[int32]int)
	for _, c := range s {
		st[c]++
	}
	for c, num := range st {
		if num < k {
			maxLen := 0
			for _, sub := range strings.Split(s, string(c)) {
				maxLen = max(maxLen, longestSubstring(sub, k))
			}
			return maxLen
		}
	}
	return len(s)
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	fmt.Println(longestSubstring("aaabb", 3))
	fmt.Println(longestSubstring("ababbc", 2))
}