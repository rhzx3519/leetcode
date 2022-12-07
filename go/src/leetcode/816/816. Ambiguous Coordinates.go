/**
@author ZhengHao Lou
Date    2022/11/07
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/ambiguous-coordinates/
*/
func ambiguousCoordinates(s string) (ans []string) {
	foo := func(x string) []string {
		n := len(x)
		if x[0] == '0' {
			if n == 1 {
				return []string{"0"}
			}
			j := 1
			for ; j < n && x[j] == '0'; j++ {
			}
			if j == n || x[n-1] == '0' {
				return []string{}
			}
			return []string{fmt.Sprintf("0.%v", x[1:])}
		}

		if x[n-1] == '0' {
			return []string{x}
		}
		var ans = []string{x}
		for i := 1; i < len(x); i++ {
			ans = append(ans, fmt.Sprintf("%v.%v", x[:i], x[i:]))
		}
		return ans
	}
	ss := s[1 : len(s)-1]
	for i := 1; i < len(ss); i++ {
		a, b := ss[:i], ss[i:]
		fmt.Println(a, b)
		a1, b1 := foo(a), foo(b)
		for _, s1 := range a1 {
			for _, s2 := range b1 {
				ans = append(ans, fmt.Sprintf("(%v, %v)", s1, s2))
			}
		}
	}
	fmt.Println(ans)
	return
}

func main() {
	//ambiguousCoordinates("(123)")
	//ambiguousCoordinates("(00011)")
	ambiguousCoordinates("(0123)")
	ambiguousCoordinates("(100)")
}
