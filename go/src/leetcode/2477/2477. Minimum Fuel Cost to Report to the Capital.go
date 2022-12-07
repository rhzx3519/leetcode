/*
*
@author ZhengHao Lou
Date    2022/11/21
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-fuel-cost-to-report-to-the-capital/
*/
func minimumFuelCost(roads [][]int, seats int) (tot int64) {
	n := len(roads) + 1
	g := make([][]int, n)
	for i := range roads {
		a, b := roads[i][0], roads[i][1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	var dfs func(pi, i int) int
	dfs = func(pi, i int) int {
		var s = 1
		for _, ni := range g[i] {
			if ni != pi {
				s += dfs(i, ni)
			}
		}
		if i != 0 {
			tot += int64((s + seats - 1) / seats)
		}
		return s
	}

	dfs(-1, 0)
	fmt.Println(tot)
	return
}

func main() {
	minimumFuelCost([][]int{{0, 1}, {0, 2}, {0, 3}}, 5)
	minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2)
	minimumFuelCost([][]int{}, 1)
}
