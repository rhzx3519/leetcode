/**
@author ZhengHao Lou
Date    2022/08/29
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/build-a-matrix-with-conditions/
*/
func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	var grid = make([][]int, k)
	for i := range grid {
		grid[i] = make([]int, k)
	}

	//n, m := len(rowConditions), len(colConditions)

	var topo func(conditions [][]int, coord map[int]int) bool
	topo = func(conditions [][]int, coord map[int]int) bool {
		ind := make([]int, k+1)
		edges := make([][]int, k+1)
		for _, r := range conditions {
			ind[r[1]]++
			edges[r[0]] = append(edges[r[0]], r[1])
		}

		var que []int
		for i := 1; i <= k; i++ {
			if ind[i] == 0 {
				que = append(que, i)
			}
		}

		for r := 0; len(que) != 0; r++ {
			i := que[0]
			que = que[1:]
			coord[i] = r
			for _, ni := range edges[i] {
				ind[ni]--
				if ind[ni] == 0 {
					que = append(que, ni)
				}
			}
		}

		return len(coord) == k
	}

	var rows, cols = make(map[int]int), make(map[int]int)
	if ok := topo(rowConditions, rows); !ok {
		return [][]int{}
	}
	if ok := topo(colConditions, cols); !ok {
		return [][]int{}
	}

	for i := 1; i <= k; i++ {
		grid[rows[i]][cols[i]] = i
	}
	fmt.Println(rows, cols)
	fmt.Println(grid)

	return grid
}

func main() {
	buildMatrix(3, [][]int{{1, 2}, {3, 2}}, [][]int{{2, 1}, {3, 2}})
	buildMatrix(3, [][]int{{1, 2}, {2, 3}, {3, 1}}, [][]int{{2, 1}})
}
