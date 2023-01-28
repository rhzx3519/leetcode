/**
@author ZhengHao Lou
Date    2023/01/08
*/
package main

import "sort"

/**
https://leetcode.cn/problems/maximize-the-minimum-powered-city/description/
思路：二分 + 差分
*/
func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	sums := make([]int, n+1)
	for i := range stations {
		sums[i+1] = sums[i] + stations[i]
	}
	return int64(sort.Search(sums[n]+k, func(minPower int) bool {
		minPower++
		diff := make([]int, n+1)
		var sumD, need int
		for i, d := range diff[:n] {
			sumD += d // 累加差分
			power := sums[min(i+r+1, n)] - sums[max(i-r, 0)]
			m := minPower - power - sumD // 当前达到minPower还需要的station数量
			if m > 0 {
				need += m
				if need > k {
					return true
				}
				sumD += m // 差分更新
				diff[min(i+r*2+1, n)] -= m
			}
		}
		return false
	}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
