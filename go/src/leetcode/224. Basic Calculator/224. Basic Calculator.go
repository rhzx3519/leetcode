package main

import "fmt"

func calculate(s string) int {
	st := []int{}
	sign := 1
	ans := 0
	n := len(s)
	for i := 0; i < n; i++ {
		c := s[i]
		if c == '(' {
			st = append(st, ans)
			ans = 0
			st = append(st, sign)
			sign = 1
		} else if c == ')' {
			ans = st[len(st)-1]*ans + st[len(st)-2]
			st = st[:len(st)-2]
		} else if c == '+' {
			sign = 1
		} else if c == '-' {
			sign = -1
		} else if isDigit(c) {
			cur := 0
			for ;i < n && isDigit(s[i]); i++ {
				cur = cur*10 + int(s[i] - '0')
			}
			i--
			ans += sign*cur
		}
	}

	return ans
}

func isDigit(c uint8) bool {
	return c >= '0'	 && c <= '9'
}


func main() {
	//fmt.Println(calculate("1 + 1"))
	//fmt.Println(calculate(" 2-1 + 2 "))
	//fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("2147483647"))
}
