/**
@author ZhengHao Lou
Date    2021/11/13
*/
package main

import "fmt"

func detectCapitalUse(word string) bool {
	n := len(word)
	if n == 1 {
		return true
	}
	if isLower(word[0]) {
		for i := 1; i < n; i++ {
			if !isLower(word[i]) {
				return false
			}
		}
	}
	if isLower(word[1]) {
		for i := 1; i < n; i++ {
			if !isLower(word[i]) {
				return false
			}
		}
		return true
	}
	for i := 1; i < n; i++ {
		if !isUpper(word[i]) {
			return false
		}
	}
	return true
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func main() {
	fmt.Println(detectCapitalUse("ggg"))
}