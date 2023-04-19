/**
@author ZhengHao Lou
Date    2021/11/05
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/remove-k-digits/
思路：单调栈
 */
func removeKdigits(num string, k int) string {
	st := []byte{}
	for i := range num {
		c := num[i]
		for k > 0 && len(st) > 0 && st[len(st) - 1] > c {
			k--
			st = st[:len(st) - 1]
		}
		st = append(st, c)
	}
	st = st[:len(st) - k]
	s := strings.TrimLeft(string(st), "0")
	if s == "" {
		return "0"
	}
	return s
}

func main() {
	fmt.Println(removeKdigits("1432219", 3))
}
