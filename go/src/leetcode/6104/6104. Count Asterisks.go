/**
@author ZhengHao Lou
Date    2022/06/27
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-asterisks/
*/
func countAsterisks(s string) int {
	var c, ans int
	for i := range s {
		switch s[i] {
		case '|':
			c++
		case '*':
			if c&1 == 0 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countAsterisks("l|*e*et|c**o|*de|"))
}
