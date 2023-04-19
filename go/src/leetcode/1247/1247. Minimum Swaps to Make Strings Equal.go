package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-swaps-to-make-strings-equal/
*/
func minimumSwap(s1 string, s2 string) (tot int) {
	var a, b int // a: x/y, b: y/x
	for i := range s1 {
		if s1[i] != s2[i] {
			if s1[i] == 'x' {
				a++
			} else {
				b++
			}
		}
	}
	tot += a >> 1
	tot += b >> 1
	a = a & 1
	b = b & 1
	if a != b {
		return -1
	}
	tot += a + b
	return
}

func main() {
	fmt.Println(minimumSwap("xx", "yy"))
	fmt.Println(minimumSwap("xy", "yx"))
	fmt.Println(minimumSwap("xx", "xy"))
	fmt.Println(minimumSwap("xxyyxyxyxx", "xyyxyxxxyx"))
}
