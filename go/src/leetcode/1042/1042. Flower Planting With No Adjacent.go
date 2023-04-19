package main

import "math/bits"

/*
*
https://leetcode.cn/problems/flower-planting-with-no-adjacent/description/
*/
func gardenNoAdj(n int, paths [][]int) []int {
	g := make([][]int, n)
	for _, p := range paths {
		a, b := p[0]-1, p[1]-1
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	color := make([]int, n)
	for i, nodes := range g {
		mask := uint8(1)
		for _, j := range nodes {
			mask |= 1 << color[j]
		}
		color[i] = bits.TrailingZeros8(^mask) // 计算最低位开始连续0的数量
	}
	return color
}
