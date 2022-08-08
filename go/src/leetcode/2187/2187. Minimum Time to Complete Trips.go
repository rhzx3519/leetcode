/**
@author ZhengHao Lou
Date    2022/03/03
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/minimum-time-to-complete-trips/
*/
func minimumTime(time []int, totalTrips int) int64 {
	var total = int64(totalTrips)
	sort.Ints(time)
	var l, r = int64(1), int64(time[len(time)-1])*total + 1
	var m int64
	for l < r {
		m = l + (r-l)>>1
		var s int64
		for _, t := range time {
			s += m / int64(t)
		}
		if s >= total {
			r = m
		} else {
			l = m + 1
		}
	}

	fmt.Println(l, r, m)
	return l
}

func main() {
	minimumTime([]int{1, 2, 3}, 5)
	minimumTime([]int{2}, 1)
	minimumTime([]int{5, 10, 10}, 9)
}
