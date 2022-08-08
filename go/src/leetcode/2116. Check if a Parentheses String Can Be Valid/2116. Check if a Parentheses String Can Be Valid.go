/**
@author ZhengHao Lou
Date    2022/01/04
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/check-if-a-parentheses-string-can-be-valid/
思路：分别记录左括号、右括号和万能符的数量，左右各遍历一次，
1. 从左往右遍历时，右括号数量不能大于左括号+万能符的数量
2. 从右往左遍历时，左括号的数量不能大于右括号+万能符的数量
同时满足1、2的才是 有效字符串
 */
func canBeValid(s string, locked string) bool {
	var n = len(s)
	if n%2 != 0 {
		return false
	}
	var l, r, m int
	for i := range s {
		if locked[i] == '1' {
			switch s[i] {
			case '(':
				l++
			case ')':
				r++
			}
		} else {
			m++
		}

		if r > l + m {
			return false
		}
	}

	l, r, m = 0, 0, 0
	for i := n-1; i >= 0; i-- {
		if locked[i] == '1' {
			switch s[i] {
			case '(':
				l++
			case ')':
				r++
			}
		} else {
			m++
		}

		if l > r + m {
			return false
		}
	}

	return true
}

func main() {
	//fmt.Println(canBeValid("))()))", "010100"))
	//fmt.Println(canBeValid("()()", "0000"))
	fmt.Println(canBeValid(")", "0"))
	//fmt.Println(canBeValid(")(", "00"))
	//fmt.Println(canBeValid("())(()(()(())()())(())((())(()())((())))))(((((((())(()))))(",
	//	"100011110110011011010111100111011101111110000101001101001111"))
	//
	//fmt.Println(canBeValid("())()))()(()(((())(()()))))((((()())(())",
	//	"1011101100010001001011000000110010100101"))	// true




}