/**
@author ZhengHao Lou
Date    2021/10/30
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/contest/biweekly-contest-64/problems/two-best-non-overlapping-events/
 */
func maxTwoEvents(events [][]int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	maximum := make([]int, n)
	var t int
	for i := n - 1; i >= 0; i-- {
		maximum[i] = max(events[i][2], t)
		t = max(events[i][2], t)
	}

	next := make([]int, n)
	for i := 0; i < n; i++ {
		j := upperBound(events, events[i][1])
		next[i] = j
	}

	fmt.Println(next, maximum)
	var ans int
	for i := 0; i < n; i++ {
		var t = events[i][2]
		if next[i] != n {
			t += maximum[next[i]]
		}
		ans= max(ans, t)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func upperBound(events [][]int, x int) int {
	l, r := 0, len(events)
	var m int
	for l < r {
		m = l + (r - l)>>1
		if events[m][0] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(maxTwoEvents([][]int{{1,3,2},{4,5,2},{2,4,3}}))
	fmt.Println(maxTwoEvents([][]int{{1,3,2},{4,5,2},{1,5,5}}))
	fmt.Println(maxTwoEvents([][]int{{1,5,3},{1,5,1},{6,6,5}}))
}