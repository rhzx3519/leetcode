/**
@author ZhengHao Lou
Date    2022/05/28
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/remove-outermost-parentheses/
*/
func removeOuterParentheses(s string) string {
	var bytes []byte
	var p, j int
	for i := range s {
		if p == 0 {
			j = i
		}
		if s[i] == '(' {
			p++
		} else {
			p--
		}
		if p == 0 {
			bytes = append(bytes, []byte(s[j+1:i])...)
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}
