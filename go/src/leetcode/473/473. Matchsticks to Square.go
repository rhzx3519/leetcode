/**
@author ZhengHao Lou
Date    2022/06/01
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/matchsticks-to-square/
思路：dfs
*/
func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	var s int
	for i := range matchsticks {
		s += matchsticks[i]
	}
	if s%4 != 0 {
		return false
	}
	var es = s / 4

	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	var dfs func(idx int, edges [4]int) bool
	dfs = func(idx int, edges [4]int) bool {
		if idx == n {
			for i := range edges {
				if edges[i] != es {
					return false
				}
			}
			return true
		}
		for i := range edges {
			if edges[i]+matchsticks[idx] <= es {
				edges[i] += matchsticks[idx]
				if dfs(idx+1, edges) {
					return true
				}
				edges[i] -= matchsticks[idx]
			}
		}
		return false
	}

	return dfs(0, [4]int{})
}

func main() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
}
