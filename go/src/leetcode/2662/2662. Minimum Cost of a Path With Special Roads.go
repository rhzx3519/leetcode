package main

import "math"

/*
*
https://leetcode.cn/problems/minimum-cost-of-a-path-with-special-roads/
*/
func minimumCost(start []int, target []int, specialRoads [][]int) int {
	type pair struct {
		x, y int
	}
	t := pair{target[0], target[1]}
	s := pair{start[0], start[1]}
	dis := make(map[pair]int, len(specialRoads)+2)
	dis[t] = math.MaxInt32
	dis[s] = 0
	vis := make(map[pair]bool, len(specialRoads)+1)
	for {
		v, dv := pair{}, -1
		for p, d := range dis {
			if !vis[p] && (dv == -1 || d < dv) {
				v = p
				dv = d
			}
		}
		if v == t {
			return dv
		}
		vis[v] = true
		dis[t] = min(dis[t], dv+abs(t.x-v.x)+abs(t.y-v.y))

		for _, sp := range specialRoads {
			w := pair{sp[2], sp[3]}
			d := dv + abs(sp[0]-v.x) + abs(sp[1]-v.y) + sp[4]
			if dw, ok := dis[w]; !ok || d < dw {
				dis[w] = d
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
