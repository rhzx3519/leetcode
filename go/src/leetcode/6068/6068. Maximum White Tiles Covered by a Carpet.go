/**
@author ZhengHao Lou
Date    2022/05/15
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/
思路：以每个interval左端点为起点，计算能够覆盖的最多瓷砖数量
	二分 + 前缀和
*/
func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		if tiles[i][0] != tiles[j][0] {
			return tiles[i][0] <= tiles[j][0]
		}
		return tiles[i][1] <= tiles[j][1]
	})

	n := len(tiles)
	sums := make([]int, n+1)
	for i := range tiles {
		sums[i+1] = sums[i] + tiles[i][1] - tiles[i][0] + 1
	}

	var mx int
	for i := 0; i < n; i++ {
		x := tiles[i][0] + carpetLen - 1
		j := upperBound(tiles, x)
		var s int
		s += sums[j] - sums[i]
		if tiles[j-1][1] >= x {
			s -= tiles[j-1][1] - x
		}
		mx = max(mx, s)
	}

	fmt.Println(mx)
	return mx
}

func upperBound(nums [][]int, x int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m][0] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maximumWhiteTiles([][]int{{1, 5}, {10, 11}, {12, 18}, {20, 25}, {30, 32}}, 10)
	maximumWhiteTiles([][]int{{10, 11}, {1, 1}}, 2)
	maximumWhiteTiles([][]int{{10, 11}, {12, 18}, {19, 20}}, 10)
}
