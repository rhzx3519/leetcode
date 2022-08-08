/**
@author ZhengHao Lou
Date    2022/06/13
*/
package main

import (
	"sort"
)

/**
https://leetcode.cn/problems/height-checker/
*/
func heightChecker(heights []int) int {
	sorted := append([]int{}, heights...)
	sort.Ints(sorted)
	var c int
	for i := range heights {
		if heights[i] != sorted[i] {
			c++
		}
	}
	return c
}

func main() {
	heightChecker([]int{1, 1, 4, 2, 1, 3})
}
