package main

import (
	"fmt"
	"strings"
)

func solveEquation(equation string) string {
	es := strings.Split(equation, "=")
	l, r := es[0], es[1]
	a1, b1 := parse(l)
	a2, b2 := parse(r)
	if a1 == a2 {
		if b1 != b2 {
			return "No solution"
		}
		return "Infinite solutions"
	}

	//fmt.Println(a1, b1, a2, b2)
	var s = fmt.Sprintf("x=%d", (b2-b1)/(a1-a2))
	fmt.Println(s)
	return s
}

func parse(s string) (a, b int) {
	hasK := false
	var k int
	sign := 1
	n := len(s)
	var i int
	for i < n {
		c := s[i]
		if isDigit(c) {
			hasK = true
			j := i
			for ; j < n && isDigit(s[j]); j++ {
				k = k*10 + int(s[j] - '0')
			}
			i = j
		} else if c == '+' {
			if hasK {
				b += sign*k
				k = 0
				hasK = false
			}
			sign = 1
			i++
		} else if c == '-' {
			if hasK {
				b += sign*k
				k = 0
				hasK = false
			}
			k = 0
			sign = -1
			i++
		}  else if c == 'x' {
			if hasK {
				a += sign*k
			} else {
				a += sign
			}
			hasK = false
			k = 0
			sign = 1
			i++
		}
	}
	if hasK {
		b += sign*k
	}
	return
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func main() {
	solveEquation("x+5-3+x=6+x-2")
	solveEquation("x=x")
	solveEquation("2x=x")
	solveEquation("2x+3x-6x=x+2")
	solveEquation("x=x+2")

	solveEquation("x=100")
	solveEquation("0x=0")
	solveEquation("-x=-1")
}
