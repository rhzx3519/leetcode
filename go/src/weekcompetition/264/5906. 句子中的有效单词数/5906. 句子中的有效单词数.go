/**
@author ZhengHao Lou
Date    2021/10/24
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/contest/weekly-contest-264/problems/number-of-valid-words-in-a-sentence/
 */
func countValidWords(sentence string) int {
	var count int
	for _, s := range strings.Split(sentence, " ") {
		if s == " " || s == "" {
			continue
		}
		if isWord(s) {
			count++
		}
	}
	return count
}

func isWord(s string) bool {
	var x, y int
	for i, c := range s {
		if c >= '0' && c <= '9' {
			return false
		}
		if c == '-' {
			if x >= 1 {
				return false
			}
			x++
			if i == 0 || i == len(s) - 1 {
				return false
			}
			if !isLow(s[i-1]) || !isLow(s[i+1]) {
				return false
			}
		}
		if c == '.' || c == '!' || c == ',' {
			if y >= 1 {
				return false
			}
			y++
			if i != len(s) - 1 {
				return false
			}
		}
	}
	return true
}

func isLow(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func main() {
	fmt.Println(countValidWords("cat and  dog"))
	//fmt.Println(countValidWords())
	//fmt.Println(countValidWords())
	//fmt.Println(countValidWords())
}