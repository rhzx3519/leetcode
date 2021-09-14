package main

import "fmt"

func longestBeautifulSubstring(word string) int {
	var seq = []byte{}
	var l, maxLen int
	for i := range word {
		if len(seq) == 0 || seq[len(seq)-1] < word[i] {
			seq = append(seq, word[i])
			l++
		} else if seq[len(seq)-1] == word[i] {
			l++
		} else {
			seq = []byte{word[i]}
			l = 1
		}

		if len(seq) == 5 {
			maxLen = max(maxLen, l)
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu"))
	fmt.Println(longestBeautifulSubstring("aeeeiiiioooauuuaeiou"))
	fmt.Println(longestBeautifulSubstring("a"))
}