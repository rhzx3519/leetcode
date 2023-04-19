package main

import (
	"fmt"
	"math/bits"
)

/*
*
https://leetcode.cn/problems/count-subtrees-with-max-distance-between-cities/description/
*/
func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]-1] = append(g[e[0]-1], e[1]-1)
		g[e[1]-1] = append(g[e[1]-1], e[0]-1)
	}

	var ans = make([]int, n-1)
	for mask := 1<<n - 1; mask != 0; mask-- {
		if mask&(mask-1) == 0 { // 至少需要两个点
			continue
		}

		var diamter, vis int
		var dfs func(x int) int
		// 求树的直径
		dfs = func(x int) (maxDep int) {
			vis |= 1 << x // 标记访问
			for _, y := range g[x] {
				if mask>>y&1 == 0 || vis>>y&1 > 0 {
					continue
				}
				depth := dfs(y) + 1
				diamter = max(diamter, maxDep+depth)
				maxDep = max(maxDep, depth)
			}
			return
		}

		dfs(bits.TrailingZeros(uint(mask)))
		if vis == mask {
			ans[diamter-1]++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(countSubgraphsForEachDiameter(4, [][]int{{1, 2}, {2, 3}, {2, 4}}))
}
