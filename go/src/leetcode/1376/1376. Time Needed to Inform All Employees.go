package main

import "fmt"

func numOfMinutes(n int, headID int, manager []int, informTime []int) (tot int) {
	g := make([][]int, n)
	for i, pi := range manager {
		if pi != -1 {
			g[pi] = append(g[pi], i)
		}
	}
	type pair struct {
		i, t int
	}
	que := []pair{{headID, informTime[headID]}}
	for len(que) != 0 {
		p := que[0]
		que = que[1:]
		tot = max(tot, p.t)
		for _, ni := range g[p.i] {
			que = append(que, pair{ni, p.t + informTime[ni]})
		}
	}

	fmt.Println(tot)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	numOfMinutes(1, 0, []int{-1}, []int{0})
	numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0})
}
