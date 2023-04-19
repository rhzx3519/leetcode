package main

/*
*
https://leetcode.cn/problems/count-number-of-possible-root-nodes/
思路：换根DP
https://leetcode.cn/problems/count-number-of-possible-root-nodes/solutions/2147714/huan-gen-dppythonjavacgo-by-endlesscheng-ccwy/
*/
func rootCount(edges [][]int, guesses [][]int, k int) (tot int) {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	type pair struct {
		x, y int
	}
	has := make(map[pair]bool)
	for _, guess := range guesses {
		has[pair{guess[0], guess[1]}] = true
	}

	var cnt0 int // 统计以0为根的正确猜测数目
	var dfs func(int, int)
	dfs = func(x int, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if has[pair{x, y}] {
					cnt0++
				}
				dfs(y, x)
			}
		}
	}

	dfs(0, -1)

	var reroot func(int, int, int)
	reroot = func(x int, fa int, c int) {
		if c >= k { // 以x为根的正确猜测数目>=k
			tot++
		}
		for _, y := range g[x] {
			if y != fa {
				cnt := c
				if has[pair{x, y}] {
					cnt--
				}
				if has[pair{y, x}] {
					cnt++
				}
				reroot(y, x, cnt)
			}
		}
	}
	// 依然以0作为根节点遍历树
	reroot(0, -1, cnt0)
	return
}
