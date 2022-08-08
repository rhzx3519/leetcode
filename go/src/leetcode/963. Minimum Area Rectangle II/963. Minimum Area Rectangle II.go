/**
@author ZhengHao Lou
Date    2021/11/25
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/minimum-area-rectangle-ii/
思路：几何题
矩形四个点如果p2和p3是对角，则p4 = p2 + p3 - p1

*/
func minAreaFreeRect(points [][]int) float64 {
	set := make(map[[2]int]bool)
	for _, p := range points {
		set[[2]int{p[0], p[1]}] = true
	}
	n := len(points)
	var ans = math.MaxFloat64
	for i := range points {
		p1 := points[i]
		for j := range points {
			if i == j {
				continue
			}
			p2 := points[j]
			for k := j+1; k < n; k++ {
				if k == i {
					continue
				}
				p3 := points[k]
				p4 := []int{p2[0] + p3[0] - p1[0], p2[1] + p3[1] - p1[1]}
				if !set[[2]int{p4[0], p4[1]}] {
					continue
				}
				f := (p3[0] - p1[0]) * (p2[0] - p1[0]) + (p3[1] - p1[1]) * (p2[1] - p1[1])
				if f != 0 {
					continue
				}
				l1 := math.Sqrt(float64((p1[0] - p3[0])*(p1[0] - p3[0]) + (p1[1] - p3[1])*(p1[1] - p3[1])))
				l2 := math.Sqrt(float64((p1[0] - p2[0])*(p1[0] - p2[0]) + (p1[1] - p2[1])*(p1[1] - p2[1])))
				s := l1*l2
				fmt.Println(p1, p2, p3, p4)
				if s < ans {
					ans = s
				}
			}
		}
	}
	if ans == math.MaxFloat64 {
		return 0
	}
	return ans
}

func main() {
	//fmt.Println(minAreaFreeRect([][]int{{1,2},{2,1},{1,0},{0,1}}))
	//fmt.Println(minAreaFreeRect([][]int{{0,1},{2,1},{1,1},{1,0},{2,0}}))
	fmt.Println(minAreaFreeRect([][]int{{0,3},{1,2},{3,1},{1,3},{2,1}}))
}
