/**
@author ZhengHao Lou
Date    2023/01/21
*/
package main

import "math"

/**
https://leetcode.cn/problems/minimum-sideway-jumps/description/
*/
const INF int = math.MaxInt32 >> 1

func minSideJumps(obstacles []int) int {
	n := len(obstacles) - 1
	f := []int{1, 0, 1}
	
	for i := 1; i <= n; i++ {
		x := obstacles[i]
		minCnt := INF
		for j := range f {
			if j != x-1 {
				minCnt = min(minCnt, f[j])
			} else {
				f[j] = INF
			}
		}
		
		for j := range f {
			if j != x-1 {
				f[j] = min(f[j], minCnt+1)
			}
		}
	}
	return min(f[0], min(f[1], f[2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
