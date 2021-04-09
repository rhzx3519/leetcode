package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	m, n := len(A), len(A[0])
	for r := 0; r < m; r++ {
		i, j := 0, n-1
		for ; i < j; i, j = i+1, j-1 {
			A[r][i], A[r][j] = A[r][j], A[r][i]
			A[r][i] ^= 1
			A[r][j] ^= 1
		}
		if i == j {
			A[r][i] ^= 1
		}
	}
	return A
}

func main() {
	A := [][]int {{1,1,0},{1,0,1},{0,0,0}}
	A = flipAndInvertImage(A)
	fmt.Println(A)
}