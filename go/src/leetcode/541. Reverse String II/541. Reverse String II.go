package main

import "fmt"

func reverseStr(s string, k int) string {
	n := len(s)
	k2 := 2*k
	bytes := []byte(s)
	for i := 0; i < n; i += k2 {
		reverse(bytes[i:min(i+k, n)])
	}
	return string(bytes)
}

func reverse(bytes []byte) {
	l, r := 0, len(bytes) - 1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcd", 2))
}