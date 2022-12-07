/**
@author ZhengHao Lou
Date    2022/09/03
*/
package main

import "fmt"

func maximumRows(mat [][]int, cols int) int {
	m, _ := len(mat), len(mat[0])
	var rows = make([]int, m)
	var c = make([]int, m)
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j]&1 != 0 {
				rows[i] |= 1 << j
				c[i]++
			}
		}
	}

	var counter = make([]int, m)
	for i := range mat {
		counter[i] = 1
		for j := range mat {
			if i == j {
				continue
			}
			if rows[i]|rows[j] == rows[i] {
				counter[i]++
			}
		}
	}

	var mx int
	for i := range mat {
		if c[i] > cols {
			continue
		}
		if counter[i] > mx {
			mx = counter[i]
		}
	}
	return mx
}

func main() {
	//fmt.Println(maximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2))
	//fmt.Println(maximumRows([][]int{{1}, {0}}, 1))
	fmt.Println(maximumRows([][]int{{0, 1}, {1, 0}}, 2))
}
