/**
@author ZhengHao Lou
Date    2022/05/03
*/
package main

import "math"

/**
https://leetcode-cn.com/problems/minimum-consecutive-cards-to-pick-up/
*/
func minimumCardPickup(cards []int) int {
	var d = math.MaxInt32
	counter := make(map[int]int)
	for i, c := range cards {
		if _, ok := counter[c]; ok && i-counter[c]+1 < d {
			d = i - counter[c] + 1
		}
		counter[c] = i
	}
	if d == math.MaxInt32 {
		return -1
	}
	return d
}
