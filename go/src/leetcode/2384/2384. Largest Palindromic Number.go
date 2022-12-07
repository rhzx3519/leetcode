/**
@author ZhengHao Lou
Date    2022/08/22
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/largest-palindromic-number/
*/
const N = 10

func largestPalindromic(num string) string {
	digits := make([]int, N)
	for i := range num {
		digits[int(num[i]-'0')]++
	}

	var bytes []byte
	var m = -1
	for i := N - 1; i >= 0; i-- {
		c := digits[i]

		if c&1 != 0 && m == -1 {
			m = i
		}

		for j := 0; j < c>>1; j++ {
			bytes = append(bytes, byte('0'+i))
		}
	}

	var hasZero bool
	var i int
	for i < len(bytes) && bytes[i] == '0' {
		hasZero = true
		i++
	}
	bytes = bytes[i:]

	var tmp []byte
	for i := len(bytes) - 1; i >= 0; i-- {
		tmp = append(tmp, bytes[i])
	}
	if m != -1 {
		bytes = append(bytes, byte('0'+m))
	}
	bytes = append(bytes, tmp...)

	if len(bytes) == 0 && hasZero {
		bytes = []byte{'0'}
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	largestPalindromic("444947137")
	largestPalindromic("00009")
}
