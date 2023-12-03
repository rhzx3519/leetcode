package main

import (
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/?envType=daily-question&envId=2023-12-03
思路：滑动窗口
*/
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	var l, s, sum int
	var mi = math.MaxInt32
	for r := range cardPoints {
		sum += cardPoints[r]
		s += cardPoints[r]
		if r-l+1 > n-k {
			s -= cardPoints[l]
			l++
		}
		if r-l+1 == n-k {
			mi = min(mi, s)
		}
	}
	return sum - mi
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	fmt.Println(maxScore([]int{2, 2, 2}, 2))
}
