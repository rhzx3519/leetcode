/**
@author ZhengHao Lou
Date    2022/06/12
*/
package main

import "unicode"

/**
https://leetcode.cn/problems/strong-password-checker-ii/
*/
func strongPasswordCheckerII(password string) bool {
	n := len(password)
	if n < 8 {
		return false
	}

	mp := make(map[rune]bool)
	for _, c := range "!@#$%^&*()-+" {
		mp[c] = true
	}

	var digit, lower, upper, special bool
	for i, c := range password {
		if unicode.IsDigit(c) {
			digit = true
		}
		if unicode.IsLower(c) {
			lower = true
		}
		if unicode.IsUpper(c) {
			upper = true
		}
		if mp[c] {
			special = true
		}
		if i > 0 && password[i-1] == password[i] {
			return false
		}
	}

	return digit && lower && upper && special
}
