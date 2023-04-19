package main

/*
*
https://leetcode.cn/problems/count-the-number-of-vowel-strings-in-range/
*/
var vowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func vowelStrings(words []string, left int, right int) int {
	n := len(words)
	sums := make([]int, n+1)
	for i := range words {
		sums[i+1] = sums[i] + vowel(words[i])
	}
	return sums[right+1] - sums[left]
}

func vowel(w string) int {
	if vowels[w[0]] && vowels[w[len(w)-1]] {
		return 1
	}
	return 0
}
