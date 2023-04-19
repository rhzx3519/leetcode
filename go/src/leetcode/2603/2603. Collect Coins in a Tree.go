package main

/*
*
https://leetcode.cn/problems/collect-coins-in-a-tree/
https://leetcode.cn/problems/collect-coins-in-a-tree/solutions/2191371/tuo-bu-pai-xu-ji-lu-ru-dui-shi-jian-pyth-6uli/
思路：拓扑排序
*/
func collectTheCoins(coins []int, edges [][]int) (tot int) {
	n := len(coins)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
		deg[u]++
		deg[v]++
	}
	// 拓扑排序去掉没有金币的叶子节点
	q := make([]int, 0, n)
	for i := range deg {
		if deg[i] == 1 && coins[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, nx := range g[x] {
			deg[nx]--
			if deg[nx] == 1 && coins[nx] == 0 {
				q = append(q, nx)
			}
		}
	}

	// 拓扑排序，从有金币的叶子节点出发
	// 同时记录节点入队的时间time[x]
	time := make([]int, n)
	for i := range deg {
		if deg[i] == 1 && coins[1] == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, nx := range g[x] {
			deg[nx]--
			if deg[nx] == 1 {
				time[nx] = time[x] + 1
				q = append(q, nx)
			}
		}
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		if time[u] >= 2 && time[v] >= 2 {
			tot += 2
		}
	}
	return
}
