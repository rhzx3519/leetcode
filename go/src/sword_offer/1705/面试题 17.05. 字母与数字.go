package main

import (
	"fmt"
	"unicode"
)

/*
*
https://leetcode.cn/problems/find-longest-subarray-lcci/
(5x, 7y) -> nx==ny, (0x,2y)(1x,3y)
*/
func findLongestSubarray(array []string) []string {
	sums := make([]int, len(array)+1)
	for i := range array {
		if unicode.IsDigit(rune(array[i][0])) {
			sums[i+1] = sums[i] + 1
		} else {
			sums[i+1] = sums[i] - 1
		}
	}
	var start, maxLen int
	var cnt = map[int]int{0: 0}
	for i := 1; i < len(sums); i++ {
		if j, ok := cnt[sums[i]]; ok {
			if i-j > maxLen {
				maxLen = i - j
				start = j
			}
		} else {
			cnt[sums[i]] = i
		}
	}
	return array[start : start+maxLen]
}

func main() {
	fmt.Println(findLongestSubarray([]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}))
	fmt.Println(findLongestSubarray([]string{"A", "A"}))
}
