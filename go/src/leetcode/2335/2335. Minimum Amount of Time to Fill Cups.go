/**
@author ZhengHao Lou
Date    2022/07/11
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/minimum-amount-of-time-to-fill-cups/
*/
func fillCups(amount []int) int {
	var c int
	for amount[0] > 0 || amount[1] > 0 || amount[2] > 0 {
		sort.Ints(amount)
		amount[2]--
		amount[1]--
		c++
	}
	return c
}

func main() {
	fmt.Println(fillCups([]int{1, 4, 2}))
	fmt.Println(fillCups([]int{5, 4, 4}))
	fmt.Println(fillCups([]int{5, 0, 0}))
}
