/**
@author ZhengHao Lou
Date    2022/04/06
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/minimum-height-trees/
*/
func findMinHeightTrees(n int, edges [][]int) []int {
	ind := make([]int, n)
	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
		ind[a]++
		ind[b]++
	}
	level := make([]int, n)
	var que []int
	for i := range ind {
		if ind[i] == 1 {
			que = append(que, i)
			level[i] = 0
		} else {
			level[i] = n + 1
		}
	}

	for lv := 0; len(que) != 0; lv++ {
		sz := len(que)
		for i := 0; i < sz; i++ {
			x := que[0]
			que = que[1:]
			if level[x] < lv {
				continue
			}
			for _, nx := range adj[x] {
				ind[nx]--
				if ind[nx] == 1 {
					que = append(que, nx)
				}
			}
			level[x] = lv
		}
	}

	var ans []int
	var mx int
	for i := range level {
		if level[i] > mx {
			ans = []int{i}
			mx = level[i]
		} else if level[i] == mx {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
}
