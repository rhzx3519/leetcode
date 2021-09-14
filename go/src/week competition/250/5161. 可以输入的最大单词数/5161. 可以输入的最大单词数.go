package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	var x = bin(brokenLetters)

	var cnt int
	words := strings.Split(text, " ")
	for _, w := range words {
		var f = true
		for _, c := range w {
			if x & (1 << int(c - 'a')) != 0 {
				f = false
				break
			}
		}
		if f {
			cnt++
		}
	}
	return cnt
}

func bin(w string) int {
	var x int
	for _, c := range w {
		x |= 1<<int(c - 'a')
	}
	return x
}

func main() {
	fmt.Println(canBeTypedWords("hello world", "ad"))
}
