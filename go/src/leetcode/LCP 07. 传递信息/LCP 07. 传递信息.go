package main

import "fmt"

func numWays(n int, relation [][]int, k int) int {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0)
	}
	for _, r := range relation {
		adj[r[0]] = append(adj[r[0]], r[1])
	}

	var ans int

	var dp func(int, int)
	dp = func(idx, c int) {
		if c == 0 {
			if idx == n-1 {
				ans++
			}
			return
		}
		for _, ni := range adj[idx] {
			dp(ni, c-1)
		}
	}
	dp(0, k)
	return ans
}

func main() {
	fmt.Println(numWays(5, [][]int{{0,2},{2,1},{3,4},{2,3},{1,4},{2,0},{0,4}}, 3))
	fmt.Println(numWays(3, [][]int{{0,2},{2,1}}, 2))
}