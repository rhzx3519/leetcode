package main

import "fmt"

/**
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
84. 柱状图中最大的矩形
使用单调递增栈，栈中第i个元素对应的高度总是小于等于栈中第i-1个元素至当前欲插入元素对应的矩阵高度
 */
func largestRectangleArea(heights []int) int {
	var a = []int{0}
	a = append(a, heights...)
	a = append(a, 0)

	var ans int
	st := []int{}
	for i := range a {
		for len(st) > 0 && a[st[len(st)-1]] > a[i] {
			t := st[len(st)-1]
			st = st[:len(st)-1]
			ans	= max(ans, a[t]*(i-st[len(st)-1]-1))
		}
		st = append(st, i)
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
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
}