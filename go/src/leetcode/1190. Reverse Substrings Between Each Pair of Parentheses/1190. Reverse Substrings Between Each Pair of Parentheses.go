package main

import "fmt"

func reverseParentheses(s string) string {
	bytes := make([]byte, 0)
	var last = []int{}
	for i := range s {
		c := s[i]
		switch c {
		case '(':
			last = append(last, len(bytes))
		case ')':
			k := last[len(last)-1]
			last = last[:len(last)-1]
			reverse(bytes[k:])
		default:
			bytes = append(bytes, c)
		}
	}

	return string(bytes)
}

func reverse(bytes []byte) {
	l, r := 0, len(bytes)-1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
}

func main() {
	fmt.Println(reverseParentheses("(abcd)"))
	fmt.Println(reverseParentheses("(u(love)i)"))
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}