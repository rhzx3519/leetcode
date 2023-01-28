/*
*
@author ZhengHao Lou
Date    2022/12/10
*/
package main

import "sort"

/*
*
https://leetcode.cn/problems/maximum-height-by-stacking-cuboids/
思路：最长递增子序列
*/
func maxHeight(cuboids [][]int) (ans int) {
	n := len(cuboids)
	for _, c := range cuboids {
		sort.Ints(c)
	}
	sort.Slice(cuboids, func(i, j int) bool {
		a, b := cuboids[i], cuboids[j]
		return a[0] < b[0] || a[0] == b[0] && (a[1] < b[1] || a[1] == b[1] && a[2] < b[2])
	})

	f := make([]int, n+1)
	for i, c2 := range cuboids {
		for j, c1 := range cuboids[:i] {
			if c1[1] <= c2[1] && c1[2] <= c2[2] { // 排序后，c1[0] <= c2[0] 恒成立
				f[i] = max(f[i], f[j]) // c1 可以堆在c2上
			}
		}
		f[i] += c2[2]
		ans = max(ans, f[i])
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
