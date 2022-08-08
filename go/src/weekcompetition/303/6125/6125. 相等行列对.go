/**
@author ZhengHao Lou
Date    2022/07/24
*/
package main

import "fmt"

func equalPairs(grid [][]int) int {
	n := len(grid)
	var rows, cols = map[string]int{}, map[string]int{}
	for i := 0; i < n; i++ {
		var bytes []byte
		for j := 0; j < n; j++ {
			bytes = append(bytes, byte(grid[i][j]))
		}
		rows[string(bytes)]++
	}

	for j := 0; j < n; j++ {
		var bytes []byte
		for i := 0; i < n; i++ {
			bytes = append(bytes, byte(grid[i][j]))
		}
		cols[string(bytes)]++
	}
	var c int
	for s := range rows {
		c += cols[s] * rows[s]
	}
	fmt.Println(c)
	return c
}

func main() {
	equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}})
	equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}})
}
