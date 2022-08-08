/**
@author ZhengHao Lou
Date    2021/11/17
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/decoded-string-at-index/
思路：
1. 先求出展开之后的字符串长度sz
2. 倒序遍历s，如果k%sz==0, 且s[i]为字母，则返回；否则，如果s[i]为字母, sz--, s[i]为数字, sz /= int(s[i])
 */

func decodeAtIndex(s string, k int) string {
	var sz int64
	var K int64 = int64(k)

	for i := range s {
		if isDigit(s[i]) {
			c := toDigit(s[i])
			sz *= c
		} else {
			sz++
		}
	}
	fmt.Println(sz)

	for i := len(s) - 1; i >= 0; i-- {
		K %= sz
		if K == 0 && !isDigit(s[i]) {
			return string(s[i])
		}

		if isDigit(s[i]) {
			sz /= toDigit(s[i])
		} else {
			sz--
		}
	}

	return ""
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func toDigit(c byte) int64 {
	return int64(c - '0')
}

func main() {
	fmt.Println(decodeAtIndex("a2b3c4d5e6f7g8h9", 9))		// "b"
	fmt.Println(decodeAtIndex("a23", 6))
	fmt.Println(decodeAtIndex("ha22", 5))
	fmt.Println(decodeAtIndex("leet2code3", 10))
	fmt.Println(decodeAtIndex("a2345678999999999999999", 1))
}
