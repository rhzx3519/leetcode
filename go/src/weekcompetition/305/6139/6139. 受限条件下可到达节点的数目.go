/**
@author ZhengHao Lou
Date    2022/08/07
*/
package main

func reachableNodes(n int, edges [][]int, restricted []int) int {
	rt := make(map[int]bool)
	for _, num := range restricted {
		rt[num] = true
	}
	adj := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	var vis = make([]bool, n)
	vis[0] = true
	que := []int{0}

	for len(que) != 0 {
		x := que[0]
		que = que[1:]
		for _, y := range adj[x] {
			if !vis[y] && !rt[y] {
				vis[y] = true
				que = append(que, y)
			}
		}
	}

	var ans int
	for i := range vis {
		if vis[i] {
			ans++
		}
	}
	return ans
}
