/**
@author ZhengHao Lou
Date    2021/11/16
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/perfect-rectangle/
思路：把每个子矩形的面积累加，四个坐标放进一个vector，然后sort一下，相同的坐标消去。
最后剩下4个出现奇数次的点，且这个四个点围成的矩形面积等于子矩形面积和，则为true
 */
func isRectangleCover(rectangles [][]int) bool {
	var s int
	mapper := make(map[[2]int]int)
	for _, rect := range rectangles {
		s += (rect[2] - rect[0]) * (rect[3] - rect[1])
		mapper[[2]int{rect[0], rect[1]}]++
		mapper[[2]int{rect[0], rect[3]}]++
		mapper[[2]int{rect[2], rect[3]}]++
		mapper[[2]int{rect[2], rect[1]}]++
	}

	left := [][2]int{}
	for k, v := range mapper {
		if v & 1 == 1 {
			left = append(left, k)
		}
	}
	if len(left) != 4 {
		return false
	}
	var x, y = math.MaxInt32 >> 1, math.MaxInt32 >> 1
	var a, b int
	for _, p := range left {
		x = min(x, p[0])
		y = min(y, p[1])
		a = max(a, p[0])
		b = max(b, p[1])
	}

	fmt.Println(mapper)
	fmt.Println(left, x, y, a, b, s)
	return (a - x) * (b - y) == s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(isRectangleCover([][]int{{1,1,3,3},{3,1,4,2},{3,2,4,4},{1,3,2,4},{2,3,3,4}}))
	//fmt.Println(isRectangleCover([][]int{{1,1,2,3},{1,3,2,4},{3,2,4,4},{3,1,4,2},{3,2,4,4}}))
	//fmt.Println(isRectangleCover([][]int{{1,1,3,3},{3,1,4,2},{1,3,2,4},{3,2,4,4}}))
	fmt.Println(isRectangleCover([][]int{{0,0,1,1},{0,1,1,2},{0,2,1,3},{1,0,2,1},{1,0,2,1},{1,2,2,3},{2,0,3,1},{2,1,3,2},{2,2,3,3}}))
}
