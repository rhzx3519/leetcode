package main

import "sort"

/*
*
https://leetcode.cn/problems/removing-minimum-number-of-magic-beans/?envType=daily-question&envId=2024-01-18
*/
func minimumRemoval(beans []int) int64 {
	n := len(beans)
	var s int64
	for i := range beans {
		s += int64(beans[i])
	}
	sort.Ints(beans)
	var tot = s
	for i, x := range beans {
		tot = min(tot, s-int64(x)*int64(n-i))
	}
	return tot
}
