/**
@author ZhengHao Lou
Date    2022/01/27
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/number-of-valid-words-in-a-sentence/
*/
func countValidWords(sentence string) int {
	var c int
	strs := strings.Split(sentence, " ")
	for _, s := range strs {
		if check(s) {
			//fmt.Println(s)
			c++
		}
	}
	fmt.Println(c)
	return c
}

func check(s string) bool {
	if s == " " || s == "" {
		return false
	}
	n := len(s)
	var c1, c2 int
	for i := range s {
		if isDigit(s[i]) {
			return false
		}
		if s[i] == '-' {
			if i == 0 || i == n-1 || !isLetter(s[i-1]) || !isLetter(s[i+1]) {
				return false
			}
			c1++
		}
		if s[i] == '!' || s[i] == '.' || s[i] == ',' {
			if i != n-1 {
				return false
			}
			c2++
		}
	}

	return c1 <= 1 && c2 <= 1
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func main() {
	countValidWords("cat and  dog")
	countValidWords("!this  1-s b8d!")
	countValidWords("alice and  bob are playing stone-game10")
	countValidWords("he bought 2 pencils, 3 erasers, and 1  pencil-sharpener.")
}
