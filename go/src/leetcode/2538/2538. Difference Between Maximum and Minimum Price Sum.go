/**
@author ZhengHao Lou
Date    2023/01/16
*/
package main

/**
https://leetcode.cn/problems/difference-between-maximum-and-minimum-price-sum/
*/
func maxOutput(n int, edges [][]int, price []int) int64 {
	var ans int
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	// 返回最大带叶子节点的路径和，最大不带叶子节点的路径和
	var dfs func(int, int) (int, int)
	dfs = func(x, fx int) (int, int) {
		p := price[x]
		var maxS1, maxS2 = p, 0
		for _, y := range g[x] {
			if y != fx {
				s1, s2 := dfs(y, x)
				ans = max(ans, max(maxS1+s2, maxS2+s1))
				maxS1 = max(maxS1, s1+p)
				maxS2 = max(maxS2, s2+p)
			}
		}
		return maxS1, maxS2
	}
	
	dfs(0, -1)
	
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
