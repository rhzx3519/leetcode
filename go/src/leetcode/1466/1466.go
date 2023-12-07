package main

/*
*
https://leetcode.cn/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/?envType=daily-question&envId=2023-12-07
*/
func minReorder(n int, connections [][]int) (tot int) {
	g := make([][][]int, n)
	for _, e := range connections {
		g[e[0]] = append(g[e[0]], []int{e[1], 1})
		g[e[1]] = append(g[e[1]], []int{e[0], 0})
	}

	var dfs func(pi, i int) int
	dfs = func(pi, i int) int {
		var res int
		for _, nxt := range g[i] {
			ni, f := nxt[0], nxt[1]
			if ni != pi {
				res += f + dfs(i, ni)
			}
		}

		return res
	}

	return dfs(-1, 0)
}
