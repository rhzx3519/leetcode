/**
@author ZhengHao Lou
Date    2022/05/03
*/
package main

/**
https://leetcode-cn.com/problems/remove-digit-from-number-to-maximize-result/
*/
func removeDigit(number string, digit byte) string {
	var s string
	for i := range number {
		if number[i] == digit {
			if number[:i]+number[i+1:] > s {
				s = number[:i] + number[i+1:]
			}
		}
	}
	return s
}
