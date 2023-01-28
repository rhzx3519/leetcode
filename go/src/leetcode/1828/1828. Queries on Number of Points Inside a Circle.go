/**
@author ZhengHao Lou
Date    2023/01/24
*/
package main

/**
https://leetcode.cn/problems/queries-on-number-of-points-inside-a-circle/
*/
func countPoints(points [][]int, queries [][]int) []int {
	var ans = make([]int, len(queries))
	for i, q := range queries {
		x, y, r := q[0], q[1], q[2]
		for _, point := range points {
			if (point[0]-x)*(point[0]-x)+(point[1]-y)*(point[1]-y) <= r*r {
				ans[i]++
			}
		}
	}
	return ans
}
