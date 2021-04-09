package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}
	mx := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		mx[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			mx[i][j] = mx[i-1][j] + mx[i][j-1] + matrix[i-1][j-1] - mx[i-1][j-1]
		}
	}
	return NumMatrix{mx}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	a := this.matrix[row2+1][col2+1]
	b := this.matrix[row1][col2+1]
	c := this.matrix[row2+1][col1]
	d := this.matrix[row1][col1]
	return a - b -c + d
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	matrix := [][]int{{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
		}
	obj := Constructor(matrix)
	for i := 1; i < len(matrix)+1; i++ {
		fmt.Println(obj.matrix[i][1:])
	}
	fmt.Println(obj.SumRegion(2, 1, 4, 3))
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
	fmt.Println(obj.SumRegion(1, 2, 2, 4))
}