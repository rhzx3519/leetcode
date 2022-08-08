/**
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/remove-covered-intervals/
 */
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] >= intervals[j][1]
	})
	n := len(intervals)
	var erase int
	var i int
	for i < n {
		var j = i + 1
		for j < n && intervals[j][1] <= intervals[i][1] {
			j++
			erase++
		}
		i = j
	}
	return n - erase
}

func main() {
	//fmt.Println(removeCoveredIntervals([][]int{{1,4},{3,6},{2,8}}))	// 2
	fmt.Println(removeCoveredIntervals([][]int{{1,2},{1,4},{3,4}}))	// 1
}
