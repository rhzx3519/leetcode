package main

/*
*
https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/description/
*/
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	cnt := make([]int, n)
	for _, trip := range trips {
		start, end := trip[0], trip[1]
		var dfs func(x, fa int) bool
		dfs = func(x, fa int) bool {
			if x == end {
				cnt[x]++
				return true
			}
			for _, nx := range g[x] {
				if nx != fa && dfs(nx, x) { // 到达终点（注意树只有唯一的一条简单路径）
					cnt[x]++
					return true
				}
			}
			return false
		}
		dfs(start, -1)
	}

	// 类似 337. 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/
	var dfs func(x, fa int) (int, int)
	dfs = func(x, fa int) (int, int) {
		notHalve := price[x] * cnt[x] // 不减半
		halve := notHalve / 2         // 减半

		for _, nx := range g[x] {
			if nx != fa {
				nh, h := dfs(nx, x)
				notHalve += min(nh, h)
				halve += nh
			}
		}

		return notHalve, halve
	}

	nh, h := dfs(0, -1)
	return min(nh, h)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
