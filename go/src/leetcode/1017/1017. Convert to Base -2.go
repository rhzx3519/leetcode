package main

import "strconv"

/**
https://leetcode.cn/problems/convert-to-base-2/description/
*/
func baseNeg2(n int) string {
	if n == 0 || n == 1 {
		return strconv.Itoa(n)
	}
	var bytes []byte
	for n != 0 {
		r := n & 1
		bytes = append(bytes, byte(r+'0'))
		n -= r
		n /= -2
	}

	for l, r := 0, len(bytes)-1; l < r; l, r = l+1, r-1 {
		bytes[l], bytes[r] = bytes[r], bytes[l]
	}
	return string(bytes)
}
