package main

import "strings"

func reversePrefix(word string, ch byte) string {
	i := strings.Index(word, string([]byte{ch}))
	if i == -1 {
		return word
	}
	bytes := []byte(word)
	l, r := 0, i
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	return string(bytes)
}
