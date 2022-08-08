/**
@author ZhengHao Lou
Date    2022/07/06
*/
package main

import (
	"fmt"
	"unicode"
)

/**
https://leetcode.cn/problems/parse-lisp-expression/
思路：
https://leetcode.cn/problems/parse-lisp-expression/solution/lisp-yu-fa-jie-xi-by-leetcode-solution-zycb/
*/
func evaluate(expression string) int {
	i, n := 0, len(expression)
	var parseVar func() string
	var parseInt func() int
	parseVar = func() string {
		i0 := i
		for i < n && expression[i] != ' ' && expression[i] != ')' {
			i++
		}
		return expression[i0:i]
	}
	parseInt = func() int {
		var sign = 1
		if expression[i] == '-' {
			sign = -1
			i++
		}
		var x int
		for i < n && unicode.IsDigit(rune(expression[i])) {
			x = x*10 + int(expression[i]-'0')
			i++
		}
		return sign * x
	}

	scope := make(map[string][]int)

	var evaluate func() int
	evaluate = func() (ret int) {
		if expression[i] != '(' { // 非表达式，可能为整数或者变量
			if unicode.IsLower(rune(expression[i])) { // 变量
				vals := scope[parseVar()]
				return vals[len(vals)-1]
			}
			return parseInt()
		}
		i++                       // 移除左括号
		if expression[i] == 'l' { // "let" 表达式
			i += 4 // 移除 "let "
			vars := []string{}
			for {
				if !unicode.IsLower(rune(expression[i])) { // 表达式最后一个变量
					ret = evaluate()
					break
				}
				vr := parseVar()
				if expression[i] == ')' {
					vals := scope[vr]
					ret = vals[len(vals)-1] // let 表达式的最后一个 expr 表达式的值
					break
				}
				vars = append(vars, vr)
				i++ // 移除空格
				scope[vr] = append(scope[vr], evaluate())
				i++ // 移除空格
			}
			for _, vr := range vars {
				scope[vr] = scope[vr][:len(scope[vr])-1] // 清除当前作用域的变量
			}
		} else if expression[i] == 'm' { // "mult" 表达式
			i += 5 // 移除 "mult"
			e1 := evaluate()
			i++ // 移除空格
			e2 := evaluate()
			ret = e1 * e2
		} else if expression[i] == 'a' { // "add" 表达式
			i += 4 // 移除 "add"
			e1 := evaluate()
			i++ // 移除空格
			e2 := evaluate()
			ret = e1 + e2
		}
		i++ // 移除右括号
		return
	}

	return evaluate()
}

func main() {
	fmt.Println(evaluate("(let x 2 (mult x (let x 3 y 4 (add x y))))"))
	fmt.Println(evaluate("(let x 3 x 2 x)"))
	fmt.Println(evaluate("(let x 1 y 2 x (add x y) (add x y))"))
}
