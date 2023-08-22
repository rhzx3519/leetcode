package main

/*
*
https://leetcode.cn/problems/maximize-distance-to-closest-person/description/
*/
func maxDistToClosest(seats []int) int {
	var maxLen, l, s, i int
	n := len(seats)
	for ; i < n && seats[i] != 1; i++ {
	}
	l = i
	for ; i < n; i++ {
		if seats[i] == 1 {
			s = 0
		} else {
			s++
		}
		maxLen = max(maxLen, s)
	}
	if maxLen&1 == 1 {
		maxLen = maxLen>>1 + 1
	} else {
		maxLen = maxLen >> 1
	}
	return max(maxLen, max(s, l))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
