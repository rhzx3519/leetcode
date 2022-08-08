/**
@author ZhengHao Lou
Date    2022/05/20
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/find-right-interval/
*/
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	index := make([]int, n)
	for i := range intervals {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		return intervals[index[i]][0] <= intervals[index[j]][0]
	})

	lowerBound := func(nums []int, x int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if intervals[nums[m]][0] >= x {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	var ans = make([]int, n)
	for i := range intervals {
		j := lowerBound(index, intervals[i][1])
		if j == n {
			ans[i] = -1
		} else {
			ans[i] = index[j]
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	findRightInterval([][]int{{1, 2}})
	findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}})
	findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}})
}
