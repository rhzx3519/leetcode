/**
@author ZhengHao Lou
Date    2022/11/15
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/maximum-units-on-a-truck/
*/
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] >= boxTypes[j][1]
	})

	var tot int
	for i := range boxTypes {
		a, b := boxTypes[i][0], boxTypes[i][1]
		c := min(truckSize, a)
		tot += c * b
		truckSize -= c
		fmt.Println(c)
	}
	fmt.Println(tot)
	return tot
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4)
	maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10)
}
