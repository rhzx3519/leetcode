/**
@author ZhengHao Lou
Date    2021/10/18
*/
package main

import "sort"

/**
https://leetcode-cn.com/problems/minimum-number-of-moves-to-seat-everyone/
 */
func minMovesToSeat(seats []int, students []int) int {
	var steps int
	n := len(seats)
	sort.Ints(seats)
	sort.Ints(students)
	for i := 0; i < n; i++ {
		steps += abs(seats[i] - students[i])
	}
	return steps
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

