/**
@author ZhengHao Lou
Date    2022/05/03
*/
package main

import (
	"sort"
	"strings"
	"unicode"
)

/**
https://leetcode-cn.com/problems/reorder-data-in-log-files/
*/
func reorderLogFiles(logs []string) []string {
	var letters, digits []string
	for _, log := range logs {
		words := strings.Split(log, " ")
		if unicode.IsLetter(rune(words[1][0])) {
			letters = append(letters, log)
		} else {
			digits = append(digits, log)
		}
	}
	
	sort.Slice(letters, func(i, j int) bool {
		w1 := strings.Split(letters[i], " ")
		w2 := strings.Split(letters[j], " ")
		c1, c2 := strings.Join(w1[1:], " "), strings.Join(w2[1:], " ")
		if c1 != c2 {
			return c1 < c2
		}
		return w1[0] < w2[0]
	})
	
	return append(letters, digits...)
}
