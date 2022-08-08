/**
@author ZhengHao Lou
Date    2022/07/01
*/
package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/**
https://leetcode.cn/problems/different-ways-to-add-parentheses/
思路：dfs
*/
func diffWaysToCompute(expression string) []int {
	var ans []int
	for i := range expression {
		if !unicode.IsDigit(rune(expression[i])) {
			left := diffWaysToCompute(expression[0:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, a := range left {
				for _, b := range right {
					var x int
					switch expression[i] {
					case '+':
						x = a + b
					case '-':
						x = a - b
					case '*':
						x = a * b
					}
					ans = append(ans, x)
				}
			}
		}
	}
	if len(ans) == 0 {
		x, _ := strconv.Atoi(expression)
		ans = append(ans, x)
	}

	return ans
}

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
