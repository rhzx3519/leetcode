/**
@author ZhengHao Lou
Date    2022/08/10
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
https://leetcode.cn/problems/solve-the-equation/
*/
func solveEquation(equation string) string {
	parse := func(s string) (a, b int) {
		n := len(s)
		var i int
		for i < n {
			var j = i
			var sign = 1

			if s[j] == '+' {
				j++
			} else if s[j] == '-' {
				sign = -1
				j++
			}
			var hasNum bool
			var num int
			for j < n && unicode.IsDigit(rune(s[j])) {
				hasNum = true
				num = num*10 + int(s[j]-'0')
				j++
			}
			if j < n && s[j] == 'x' {
				if num == 0 && !hasNum {
					a += sign
				} else {
					a += sign * num
				}
				j++
			} else {
				b += sign * num
			}
			i = j
		}
		return
	}

	strs := strings.Split(equation, "=")
	left, right := strs[0], strs[1]
	a1, b1 := parse(left)
	a2, b2 := parse(right)
	a := a1 - a2
	b := b2 - b1
	if a == 0 {
		if b != 0 {
			return "No solution"
		}
		return "Infinite solutions"
	}

	return fmt.Sprintf("x=%v", b/a)
}

func main() {
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
	fmt.Println(solveEquation("x=x"))
	fmt.Println(solveEquation("2x=x"))
	fmt.Println(solveEquation("0x=0"))
}
