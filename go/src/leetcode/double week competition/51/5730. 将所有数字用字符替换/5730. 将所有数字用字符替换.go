package main

import "fmt"

func replaceDigits(s string) string {
	runes := make([]rune, 0)
	for i, c := range s {
		if i&1 != 0 {
			runes = append(runes, shift(rune(s[i-1]), int(c-'0')))
		} else {
			runes = append(runes, c)
		}
	}
	return string(runes)
}

func shift(c rune, x int) rune {
	return rune(int(c) + x)
}

func main() {
	fmt.Println(replaceDigits("a1c1e1"))
	fmt.Println(replaceDigits("a1b2c3d4e"))
}