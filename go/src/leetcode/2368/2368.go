package main

import "fmt"

/*
*
https://leetcode.cn/problems/reachable-nodes-with-restrictions/?envType=daily-question&envId=2024-03-02
*/
func reachableNodes(n int, edges [][]int, restricted []int) (tot int) {
	rst := make(map[int]bool)
	for _, i := range restricted {
		rst[i] = true
	}
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	q := []int{0}
	for len(q) != 0 {
		r := q[0]
		q = q[1:]
		rst[r] = true
		tot++
		for _, ni := range g[r] {
			if !rst[ni] {
				q = append(q, ni)
			}
		}
	}
	return
}

func main() {
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5}))
}
