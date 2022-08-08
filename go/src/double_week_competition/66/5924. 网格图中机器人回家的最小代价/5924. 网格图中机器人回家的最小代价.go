/**
@author ZhengHao Lou
Date    2021/11/27
*/
package main

import (
	"fmt"
)
//
var dir = [][]int{{1,0},{0,-1},{-1,0},{0,1}}
func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	//m, n := len(rowCosts), len(colCosts)
	var cost int
	if startPos[0] < homePos[0] {
		for i := startPos[0] + 1; i <= homePos[0]; i++ {
			cost += rowCosts[i]
		}
	} else {
		for i := startPos[0] - 1; i >= homePos[0]; i-- {
			cost += rowCosts[i]
		}
	}

	if startPos[1] < homePos[1] {
		for i := startPos[1] + 1; i <= homePos[1]; i++ {
			cost += colCosts[i]
		}
	} else {
		for i := startPos[1] - 1; i >= homePos[1]; i-- {
			cost += colCosts[i]
		}
	}

	return cost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCost([]int{1,0}, []int{2,3}, []int{5,4,3}, []int{8,2,6,7}))
	fmt.Println(minCost([]int{0,0}, []int{0,0}, []int{5}, []int{8}))
	fmt.Println(minCost([]int{0,0}, []int{1,2}, []int{10,2,5}, []int{10,6,1}))	// 9
}