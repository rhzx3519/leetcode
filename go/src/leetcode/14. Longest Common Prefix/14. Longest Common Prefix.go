package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	t := strs[0]
	for _, str := range strs[1:] {
		j := 0
		for ; j < min(len(t), len(str)); j++ {
			if str[j] != t[j] {
				break
			}
		}
		t = t[:j]
	}
	return t
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	//fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
}