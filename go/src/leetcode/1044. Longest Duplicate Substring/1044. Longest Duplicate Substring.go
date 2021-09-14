package main

import "fmt"

func longestDupSubstring(s string) string {
	n := len(s)
	var mapper map[string]int
	for l := n-1; l >= 1; l-- {	// l is the length of sub string
		mapper = map[string]int{}
		for i := 0; i <= n - l; i++ {
			sub := s[i:i+l]
			if _, ok := mapper[sub]; ok {
				return sub
			}
			mapper[sub]	= i
		}
	}
	return ""
}

func main() {
	fmt.Println(longestDupSubstring("banana"))
	fmt.Println(longestDupSubstring(""))
	fmt.Println(longestDupSubstring("aa"))
}
