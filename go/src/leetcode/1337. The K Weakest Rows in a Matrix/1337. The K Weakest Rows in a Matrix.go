package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	fights := make([]int, 100)
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		var f int
		for j := 0; j < n && mat[i][j] == 1; j++ {
			f++
		}
		fights[i] = f
	}

	indice := make([]int, m)
	for i := 0; i < len(indice); i++ {
		indice[i] = i
	}
	sort.Slice(indice, func(i, j int) bool {
		x, y := indice[i], indice[j]
		if fights[x] != fights[y] {
			return fights[x] < fights[y]
		}
		return x < y
	})
	return indice[:k]
}
