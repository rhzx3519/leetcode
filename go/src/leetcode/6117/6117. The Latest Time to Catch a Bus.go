/**
@author ZhengHao Lou
Date    2022/07/10
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/the-latest-time-to-catch-a-bus/
思路：贪心
*/
func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	n, m := len(buses), len(passengers)
	var i, j, c int
	for ; i < n; i++ {
		c = capacity
		for j < m && passengers[j] <= buses[i] && c != 0 {
			c--
			j++
		}
	}
	j--

	var last int
	if c != 0 {
		last = buses[n-1]
	} else {
		last = passengers[j]
	}
	for j >= 0 && passengers[j] == last {
		last--
		j--
	}
	fmt.Println(last)
	return last
}

func main() {
	latestTimeCatchTheBus([]int{10, 20}, []int{2, 17, 18, 19}, 2)
	latestTimeCatchTheBus([]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2)
}
