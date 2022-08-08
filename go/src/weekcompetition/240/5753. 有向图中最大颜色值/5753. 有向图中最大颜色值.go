package main

import (
	"fmt"
)

var cache = make(map[int]int)

func largestPathValue(colors string, edges [][]int) int {
	var ans int
	cache = make(map[int]int)
	n := len(colors)
	var adj = make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}

	vis := make([]bool, n)

	cs := make([]int, 26)

	var dp func(idx int) int
	dp = func(idx int) int {
		//if i, ok := cache[idx]; ok {
		//	return i
		//}
		maxVal := 0
		for _, ni := range adj[idx] {
			if vis[ni] {
				return -1
			}
			vis[ni] = true
			cs[colors[idx] - 'a']++
			t := dp(ni)
			if maxVal < t {
				maxVal = t
			}
			cs[colors[idx] - 'a']--
			vis[ni] = false
		}
		return maxVal
	}

	for i := 0; i < n; i++ {
		vis[i] = true
		maxVal := dp(i)
		vis[i] = false
		if maxVal == -1 {
			return -1
		}
		if maxVal > ans {
			ans = maxVal
		}
	}

	return ans
}

func max(cs []int) int {
	var maxVal int
	for i := 0; i < len(cs); i++ {
		if cs[i] > maxVal {
			maxVal = cs[i]
		}
	}
	return maxVal
}

func main() {
	fmt.Println(largestPathValue("abaca", [][]int{{0,1},{0,2},{2,3},{3,4}}))
}