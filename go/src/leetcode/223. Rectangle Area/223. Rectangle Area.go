/**
@author ZhengHao Lou
Date    2021/09/30

https://leetcode-cn.com/problems/rectangle-area/

*/
package main

import "fmt"

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	var s = (ax2 - ax1) * (ay2 - ay1) + (bx2 - bx1) * (by2 - by1)
	if bx1 >= ax2 || ax1 >= bx2 {
		return s
	}
	if by1 >= ay2 || ay1 >= by2 {
		return s
	}

	var cx1, cy1, cx2, cy2 int
	// left-down
	cx1 = max(ax1, bx1)
	cy1 = max(ay1, by1)
	// right-up
	cx2 = min(bx2, ax2)
	cy2 = min(by2, ay2)
	return s - (cx2 - cx1) * (cy2 - cy1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(computeArea(-3,0,3,4,0,-1,9,2))
}