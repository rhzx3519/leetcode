/**
@author ZhengHao Lou
Date    2021/10/02
*/
package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	k := len(original)
	if m*n != k {
		return [][]int{}
	}
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := range original {
		r, c := i / n, i % n
		f[r][c] = original[i]
	}
	fmt.Println(f)
	return f
}

func main() {
	construct2DArray([]int{1,2,3,4}, 2, 2)
	construct2DArray([]int{1,2,3}, 1, 3)
	construct2DArray([]int{1,2}, 1, 1)
	construct2DArray([]int{3}, 1, 2)
}
