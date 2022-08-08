/**
@author ZhengHao Lou
Date    2022/07/05
*/
package main

import (
	"fmt"
	"unicode"
)

/**
https://leetcode.cn/problems/decode-the-message/
*/
const N int = 26

func decodeMessage(key string, message string) string {
	letters := make([]int, N)
	for i := range letters {
		letters[i] = -1
	}
	var c int
	for i := range key {
		if !unicode.IsLetter(rune(key[i])) {
			continue
		}
		j := int(key[i] - 'a')
		if letters[j] == -1 {
			letters[j] = c
			c++
		}
	}
	var bytes []byte
	for i := range message {
		switch message[i] {
		case ' ':
			bytes = append(bytes, ' ')
		default:
			bytes = append(bytes, byte('a'+letters[int(message[i]-'a')]))

		}
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv")
	decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb")
}
