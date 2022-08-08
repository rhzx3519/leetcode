/**
@author ZhengHao Lou
Date    2022/05/15
*/
package main

/**
https://leetcode.cn/problems/largest-triangle-area/
*/
func largestTriangleArea(points [][]int) float64 {
	var s float64
	for i := range points {
		for j := range points {
			for k := range points {
				s = max(s, area(points[i], points[j], points[k]))
			}
		}
	}
	return s
}

func area(a, b, c []int) float64 {
	return float64(a[0]*b[1]+b[0]*c[1]+c[0]*a[1]-a[0]*c[1]-b[0]*a[1]-c[0]*b[1]) / 2.0
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
