/**
@author ZhengHao Lou
Date    2022/01/02
*/
package main

import "fmt"

func checkString(s string) bool {
	var a, b = -1, 101
	for i := range s {
		switch s[i] {
		case 'a':
			a = i
		case 'b':
			b = i
		}
		if b < a {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkString("baaaabbb"))
	fmt.Println(checkString("abab"))
	fmt.Println(checkString("bbb"))
}
