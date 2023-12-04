package main

import "fmt"

/*
*
https://leetcode.cn/problems/separate-black-and-white-balls/
*/
func minimumSteps(s string) (tot int64) {
	var c1 int
	for i := range s {
		switch s[i] {
		case '1':
			c1++
		case '0':
			tot += int64(c1)
		}
	}
	return
}

func main() {
	fmt.Println(minimumSteps("0111"))
}
