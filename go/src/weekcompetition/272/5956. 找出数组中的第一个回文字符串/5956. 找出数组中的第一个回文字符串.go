/**
@author ZhengHao Lou
Date    2021/12/19
*/
package main

func firstPalindrome(words []string) string {
	for _, w := range words {
		if pal(w) {
			return w
		}
	}
	return ""
}

func pal(word string) bool {
	l, r := 0, len(word) - 1
	for l < r {
		if word[l] != word[r] {
			return false
		}
		l++
		r--
	}
	return true
}
