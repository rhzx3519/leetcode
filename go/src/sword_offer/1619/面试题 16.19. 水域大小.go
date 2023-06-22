package main

import (
	"fmt"
	"sort"
)

var diff = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func pondSizes(land [][]int) []int {
	m, n := len(land), len(land[0])
	var dfs func(int, int) int
	dfs = func(x int, y int) int {
		land[x][y] = 1
		var sz = 1
		for _, d := range diff {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && land[nx][ny] == 0 {
				sz += dfs(nx, ny)
			}
		}
		return sz
	}
	var tot []int
	for i := range land {
		for j := range land[i] {
			if land[i][j] == 0 {
				sz := dfs(i, j)
				tot = append(tot, sz)
			}
		}
	}
	sort.Ints(tot)
	return tot
}

func main() {
	fmt.Println(pondSizes([][]int{{0, 2, 1, 0}, {0, 1, 0, 1}, {1, 1, 0, 1}, {0, 1, 0, 1}}))
}
