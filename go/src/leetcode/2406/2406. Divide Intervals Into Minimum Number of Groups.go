/**
@author ZhengHao Lou
Date    2022/09/12
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/divide-intervals-into-minimum-number-of-groups/
*/
func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	var last []int
	for i := range intervals {
		j := lowerBound(last, intervals[i][0])
		if j > 0 && last[j-1] < intervals[i][0] {
			copy(last[j-1:], last[j:])
			last = last[:len(last)-1]
		}
		j = lowerBound(last, intervals[i][1])
		last = append(last, 0)
		copy(last[j+1:], last[j:])
		last[j] = intervals[i][1]
	}

	fmt.Println(last)
	return len(last)
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	//fmt.Println(lowerBound([]int{5, 8}, 5))
	minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}})
	minGroups([][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}})
}
