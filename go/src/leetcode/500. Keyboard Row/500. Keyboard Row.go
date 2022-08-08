/**
@author ZhengHao Lou
Date    2021/10/31
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/keyboard-row/
 */
const (
	s1 string = "qwertyuiop"
	s2 string = "asdfghjkl"
	s3 string = "zxcvbnm"
)
var (
	m1 = map[byte]int{}
	m2 = map[byte]int{}
	m3 = map[byte]int{}
)

func findWords(words []string) []string {
	for i := range s1 {
		m1[s1[i]]++
	}
	for i := range s2 {
		m2[s2[i]]++
	}
	for i := range s3 {
		m3[s3[i]]++
	}

	var check = func(w string) bool {
		w = strings.ToLower(w)
		l := getLine(w[0])
		for i := 1; i < len(w); i++ {
			if getLine(w[i]) != l {
				return false
			}
		}
		return true
	}

	ans := []string{}
	for _, word := range words {
		if check(word) {
			ans = append(ans, word)
		}
	}

	fmt.Println(ans)
	return ans
}



func getLine(c byte) int {
	if m1[c] != 0 {
		return 1
	}
	if m2[c] != 0 {
		return 2
	}
	if m3[c] != 0 {
		return 3
	}
	return -1
}

func main() {
	findWords([]string{"Hello","Alaska","Dad","Peace"})
}



