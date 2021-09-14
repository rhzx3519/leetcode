package main

import "fmt"

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	bytes := []byte(s)
	l, r := 0, len(s) - 1
	for l < r {
		for l < r {
			if _, ok := vowels[s[l]]; ok {
				break
			}
			l++
		}

		for l < r {
			if _, ok := vowels[s[r]]; ok {
				break
			}
			r--
		}

		if l < r {
			bytes[l], bytes[r] = bytes[r], bytes[l]
			l++
			r--
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}
