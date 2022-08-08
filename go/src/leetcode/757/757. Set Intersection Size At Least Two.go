/**
@author ZhengHao Lou
Date    2022/07/22
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/set-intersection-size-at-least-two/
思路：贪心，根据左边界升序排序，有边界降序排序，分类讨论


*/
func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	var c = 2
	n := len(intervals)
	var cur, next = intervals[n-1][0], intervals[n-1][0] + 1
	for i := n - 2; i >= 0; i-- {
		if intervals[i][1] >= next {
			continue
		} else if intervals[i][1] < cur {
			cur = intervals[i][0]
			next = intervals[i][0] + 1
			c += 2
		} else {
			next = cur
			cur = intervals[i][0]
			c++
		}
	}
	return c
}

func main() {
	//fmt.Println(intersectionSizeTwo([][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}))
	//fmt.Println(intersectionSizeTwo([][]int{{2, 3}, {3, 4}, {5, 10}, {5, 8}}))
	fmt.Println(intersectionSizeTwo([][]int{{0, 2}, {1, 3}, {2, 3}}))
}
