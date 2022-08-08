package main

import (
	"fmt"
	"math"
)

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	f := make([][]int, n1 + 1)
	for i := range f {
		f[i] = make([]int, n2 + 1)
		for j := range f[i] {
			f[i][j] = math.MaxInt32 >> 1
		}
	}
	f[0][0] = 0
	for i := 1; i <= n1; i++ {
		f[i][0] = i
	}
	for j := 1; j <= n2; j++ {
		f[0][j] = j
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = min(f[i][j], f[i-1][j-1])
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-1]) + 1
			}
		}
	}

	return f[n1][n2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDistance("sea", "eat"))
}
