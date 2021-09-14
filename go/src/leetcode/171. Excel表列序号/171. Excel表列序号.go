package main

import "fmt"

/**
https://leetcode-cn.com/problems/excel-sheet-column-number/
 */
func titleToNumber(columnTitle string) int {
	var s int
	for _, c := range columnTitle {
		i := int(byte(c) - 'A')
		s = s*26 + i + 1
	}
	return s
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}