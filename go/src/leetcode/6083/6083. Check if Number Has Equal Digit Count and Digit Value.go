/**
@author ZhengHao Lou
Date    2022/05/29
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/check-if-number-has-equal-digit-count-and-digit-value/
*/
func digitCount(num string) bool {
	counter := make([]int, 10)
	for i := range num {
		counter[int(num[i]-'0')]++
	}
	fmt.Println(counter)
	for i := range num {
		if int(num[i]-'0') != counter[i] {
			return false
		}
	}
	return true
}

func main() {
	digitCount("1210")
}
