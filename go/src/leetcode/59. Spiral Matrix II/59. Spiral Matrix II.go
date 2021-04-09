package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	t := 1
	li, ri, lj, rj := 0, n-1, 0, n-1
	for li <= ri && lj <= rj {
		i, j := li, lj
		for ; j <= rj; j++ {
			matrix[i][j] = t
			t++
		}
		i++
		j--
		for ; i <= ri; i++ {
			matrix[i][j] = t
			t++
		}
		i--
		j--
		for ; li != ri && j > lj; j-- {
			matrix[i][j] = t
			t++
		}
		for ; lj != rj && i > li; i-- {
			matrix[i][j] = t
			t++
		}
		li++
		lj++
		ri--
		rj--
	}

	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
