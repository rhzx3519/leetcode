/**
@author ZhengHao Lou
Date    2022/05/29
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/maximum-total-importance-of-roads/
*/
func maximumImportance(n int, roads [][]int) int64 {
	adj := make([][]int, n)
	for _, road := range roads {
		adj[road[0]] = append(adj[road[0]], road[1])
		adj[road[1]] = append(adj[road[1]], road[0])
	}

	importance := make([]int, n)
	for i := range importance {
		importance[i] = i
	}
	sort.Slice(importance, func(i, j int) bool {
		return len(adj[importance[i]]) > len(adj[importance[j]])
	})

	imp := make([]int, n)

	for i := range imp {
		imp[importance[i]] = n - i
	}

	var ans int64
	counter := make([]bool, n)
	for i := 0; i < n; i++ {
		for _, ni := range adj[i] {
			if counter[ni] {
				continue
			}
			ans += int64(imp[i] + imp[ni])
		}
		counter[i] = true
	}

	fmt.Println(ans)
	return ans
}

func main() {
	maximumImportance(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}})
}
