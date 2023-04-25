package main

import "math"

/*
*
https://leetcode.cn/problems/filling-bookcase-shelves/description/
*/
func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt32
	}
	f[0] = 0
	for i := 1; i <= n; i++ {
		var w, h int
		for j := i - 1; j >= 0; j-- {
			w += books[j][0]
			if w > shelfWidth {
				break
			}
			h = max(h, books[j][1])
			f[i] = min(f[i], f[j]+h)
		}
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
