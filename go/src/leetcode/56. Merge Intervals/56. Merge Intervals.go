package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	erase := make(map[int]int)
	n := len(intervals)
	last := intervals[0]
	for i := 1; i < n; i++ {
		if last[1] >= intervals[i][0] {
			last[1] = max(last[1], intervals[i][1])
			erase[i]++
		} else {
			last = intervals[i]
		}
	}

	ans := [][]int{}
	for i := 0; i < n; i++ {
		if erase[i] != 0 {
			continue
		}
		ans = append(ans, intervals[i])
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
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
	fmt.Println(merge([][]int{{1,4},{4,5}}))
}
