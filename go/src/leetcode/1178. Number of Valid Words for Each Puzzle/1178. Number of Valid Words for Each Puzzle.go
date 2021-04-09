package main

import (
	"fmt"
)

func findNumOfValidWords(words []string, puzzles []string) []int {
	np, nw := len(puzzles), len(words)
	ans := make([]int, np)

	hash := make(map[int]int)
	for i := 0; i < nw; i++ {
		t := 0
		for _, ch := range words[i] {
			t |= 1<<(ch-'a')
		}
		hash[t]++
	}

	for i := 0; i < np; i++ {
		k, t := 0, puzzles[i][0] - 'a'

		for _, ch := range puzzles[i] {
			k |= 1<<(ch-'a')
		}
		// 枚举k的所有组合
		for j := k; j > 0; j = (j-1)&k {
			if (j>>t & 1) == 1 {
				ans[i] += hash[j]
			}
		}
	}

	return ans
}


func main() {
	words := []string{"aaaa","asas","able","ability","actt","actor","access"}
	puzzles := []string{"aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"}
	fmt.Println(findNumOfValidWords(words, puzzles))
}