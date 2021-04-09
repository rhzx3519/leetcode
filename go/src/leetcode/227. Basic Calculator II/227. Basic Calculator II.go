package main

import "fmt"

func calculate(s string) int {
	st := make([]int, 0)
	ans := 0
	sign := '+'
	d := 0
	n := len(s)
	for i := 0; i < n; i++ {
		c := rune(s[i])
		if isDigit(c) {
			d = d * 10 + int(s[i] - '0')
		}
		if (!isDigit(c) && c != ' ') || i == n-1 {
			switch sign {
			case '+':
				st = append(st, d)
			case '-':
				st = append(st, -d)
			case '*':
				last := st[len(st)-1]
				st = st[:len(st)-1]
				st = append(st, last*d)
			case '/':
				last := st[len(st)-1]
				st = st[:len(st)-1]
				st = append(st, last/d)
			}
			sign = c
			d = 0
		}
	}

	for len(st) > 0 {
		ans += st[len(st)-1]
		st = st[:len(st)-1]
	}
	return ans
}

func isDigit(c rune) bool {
	return c >= '0'	&& c <= '9'
}

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate("3"))
	fmt.Println(calculate("3/2"))
	fmt.Println(calculate("3+5 /2"))
}
