/**
@author ZhengHao Lou
Date    2021/11/25
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/vowel-spellchecker/
 */
func spellchecker(wordlist []string, queries []string) []string {
	wordMap := make(map[string]int)
	lowerWordMap := []string{}
	vowelWordMap := []string{}
	for _, w := range wordlist {
		wordMap[w]++
		lowerWordMap = append(lowerWordMap, strings.ToLower(w))
		vowelWordMap = append(vowelWordMap, eraseVowel(w))
	}
	var ans = []string{}
	for _, q := range queries {
		if wordMap[q] != 0 {
			ans = append(ans, q)
		} else if i := match(lowerWordMap, strings.ToLower(q)); i != -1 {
			ans = append(ans, wordlist[i])
		} else if i := match(vowelWordMap, eraseVowel(q)); i != -1 {
			ans = append(ans, wordlist[i])
		} else {
			ans = append(ans, "")
		}
	}

	return ans
}

func match(wordList []string, q string) int {
	for i, w := range wordList {
		if w == q {
			return i
		}
	}
	return -1
}


func eraseVowel(w string) string {
	bytes := []byte{}
	for _, c := range strings.ToLower(w) {
		if isVowel(byte(c)) {
			bytes = append(bytes, '-')
		} else {
			bytes = append(bytes, byte(c))
		}
	}
	return string(bytes)
}

func isVowel(c byte) bool {
	switch c {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	}
	return false
}

func main() {
	fmt.Println(spellchecker([]string{"KiTe","kite","hare","Hare"},
	[]string{"kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"}))
}
