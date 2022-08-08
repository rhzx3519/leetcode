/**
@author ZhengHao Lou
Date    2022/08/01
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/longest-cycle-in-a-graph/
*/
func longestCycle(edges []int) int {
	n := len(edges)
	time := make([]int, n)
	var clock, ans = 1, -1
	for i, t := range time {
		if t > 0 {
			continue
		}

		for start := clock; i >= 0; i = edges[i] {
			if time[i] != 0 {
				if time[i] >= start {
					ans = max(ans, clock-time[i])
				}
				break
			}
			time[i] = clock
			clock++
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
	fmt.Println(longestCycle([]int{3, 3, 4, 2, 3}))
	fmt.Println(longestCycle([]int{2, -1, 3, 1}))
}
