/**
@author ZhengHao Lou
Date    2022/05/16
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/insert-interval/
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	var c int
	l, r := newInterval[0], newInterval[1]
	i := lowerBound(intervals, l) // i := lowerBound(this.intervals, l - 1), 可以合并[5,6][7,10]这种相邻的区间
	for i < len(intervals) {
		if intervals[i][0] > r { // this.intervals[i][0] > r +  1 ，可以合并[5,6][7,10]这种相邻的区间
			break
		}
		l = min(l, intervals[i][0])
		r = max(r, intervals[i][1])
		c -= intervals[i][1] - intervals[i][0] + 1
		copy(intervals[i:], intervals[i+1:])
		intervals = intervals[:len(intervals)-1]
	}
	c += r - l + 1
	intervals = append(intervals, []int{})
	copy(intervals[i+1:], intervals[i:])
	intervals[i] = []int{l, r}

	fmt.Println(c, intervals)
	return intervals
}

func lowerBound(intervals [][]int, x int) int {
	l, r := 0, len(intervals)
	for l < r {
		m := l + (r-l)>>1
		if intervals[m][1] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	insert([][]int{{1, 5}}, []int{6, 8})
	insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	insert([][]int{}, []int{5, 7})
	insert([][]int{{1, 5}}, []int{2, 3})
	insert([][]int{{1, 5}}, []int{2, 7})
}
