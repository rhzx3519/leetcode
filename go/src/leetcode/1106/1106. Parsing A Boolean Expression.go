/**
@author ZhengHao Lou
Date    2022/11/05
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/parsing-a-boolean-expression/description/
*/
func parseBoolExpr(expression string) bool {
	var st1, st2 []byte
	for i := range expression {
		if expression[i] == '!' || expression[i] == '&' || expression[i] == '|' {
			st1 = append(st1, expression[i])
		} else if expression[i] == ')' {
			var t, f int
			for len(st2) != 0 && st2[len(st2)-1] != '(' {
				if st2[len(st2)-1] == 't' {
					t++
				} else if st2[len(st2)-1] == 'f' {
					f++
				}
				st2 = st2[:len(st2)-1]
			}
			st2 = st2[:len(st2)-1] // pop '('
			op := st1[len(st1)-1]
			st1 = st1[:len(st1)-1]
			switch op {
			case '!':
				if t > 0 {
					st2 = append(st2, 'f')
				} else {
					st2 = append(st2, 't')
				}
			case '&':
				if f > 0 {
					st2 = append(st2, 'f')
				} else {
					st2 = append(st2, 't')
				}
			case '|':
				if t > 0 {
					st2 = append(st2, 't')
				} else {
					st2 = append(st2, 'f')
				}
			}

		} else {
			st2 = append(st2, expression[i])
		}
	}
	fmt.Println(st1, st2, string(st2[len(st2)-1]))
	return st2[len(st2)-1] == 't'
}

func main() {
	fmt.Println(parseBoolExpr("!(f)"))
	fmt.Println(parseBoolExpr("|(f,t)"))
	fmt.Println(parseBoolExpr("&(t,f)"))
}
