/**
@author ZhengHao Lou
Date    2023/01/22
*/
package main

import (
	"fmt"
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	type pair struct {
		id, s int
	}
	m, n := len(score), len(score[0])
	var ids = make([]pair, m)
	for i := 0; i < m; i++ {
		ids[i] = pair{id: i, s: score[i][k]}
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i].s > ids[j].s
	})
	g := make([][]int, m)
	for i := range g {
		g[i] = make([]int, n)
		copy(g[i], score[ids[i].id])
	}
	fmt.Println(g)
	return g
}

func main() {
	sortTheStudents([][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}, 2)
}
