package main

import "fmt"

/*
*
https://leetcode.cn/problems/count-vowel-strings-in-ranges/
*/
var vowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	var sums = make([]int, n+1)
	for i := range words {
		sums[i+1] = sums[i] + vowel(words[i])
	}
	var ans = make([]int, len(queries))
	for i := range queries {
		l, r := queries[i][0], queries[i][1]
		ans[i] = sums[r+1] - sums[l]
	}
	return ans
}

func vowel(w string) int {
	if vowels[w[0]] && vowels[w[len(w)-1]] {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
}
