/**
@author ZhengHao Lou
Date    2021/12/03
*/
package 最短路径算法

import (
	"fmt"
)

func ExampleSPFA_1() {
	grid := buildGrid(4, [][]int{{1,0,1},{1,2,1},{2,3,1}})
	dist := SPFA(grid, 4, 1)
	fmt.Println(dist)
	// Output:
	// [1 0 1 2]
}

func ExampleSPFA_2() {
	grid := buildGrid(2, [][]int{{0,1,1}})
	dist := SPFA(grid, 2, 1)
	fmt.Println(dist)
	// Output:
	// [1073741823 0]
}


func buildGrid(n int, edges [][]int) [][]int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
		for j := range grid[i] {
			grid[i][j] = INF
		}
	}

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		grid[u][v] = w
	}

	return grid
}