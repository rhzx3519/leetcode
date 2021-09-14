package main

import "fmt"

func findRotation(mat [][]int, target [][]int) bool {
	for k := 0; k < 4; k++ {
		if equal(mat, target) {
			return true
		}
		mat = rotate(mat)
		fmt.Println(mat)
	}

	return false
}

func rotate(mat [][]int) [][]int {
	n := len(mat)
	t := make([][]int, 0)
	
	for j := 0; j < n; j++ {
		row := []int{}
		for i := n-1; i >= 0; i-- {
			row = append(row, mat[i][j])
		}
		t = append(t, row)
	}
	return t
}

func equal(m1, m2 [][]int) bool {
	n := len(m1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(findRotation([][]int{{0,1},{1,0}}, [][]int{{1,0},{0,1}}))
	fmt.Println(findRotation([][]int{{0,1},{1,1}}, [][]int{{1,0},{0,1}}))
	fmt.Println(findRotation([][]int{{0,1},{1,1}}, [][]int{{1,0},{0,1}}))
}
