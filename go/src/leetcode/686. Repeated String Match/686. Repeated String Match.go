/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/repeated-string-match/
思路：重复之后可能包含b的最大长度为2*a+b，
 */
func repeatedStringMatch(a string, b string) int {
	var c = a
	for step := 1; len(c) < len(a)<<1 + len(b); step++ {
		if strings.Index(c, b) != -1 {
			return step
		}
		c = string(append([]byte(c), []byte(a)...))
	}
	return -1
}

func main() {
	fmt.Println(repeatedStringMatch("aa", "a"))
	//fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	//fmt.Println(repeatedStringMatch("a", "aa"))
	//fmt.Println(repeatedStringMatch("a", "a"))
	//fmt.Println(repeatedStringMatch("abc", "wxyz"))
}