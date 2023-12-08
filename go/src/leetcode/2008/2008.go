package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/maximum-earnings-from-taxi/?envType=daily-question&envId=2023-12-08
思路：动态规划
*/
func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})

	m := len(rides)
	f := make([]int64, m+1)
	for i := range rides {
		j := sort.Search(i+1, func(k int) bool {
			return rides[k][1] > rides[i][0]
		})
		f[i+1] = max(f[i], f[j]+int64(rides[i][1]-rides[i][0]+rides[i][2]))
	}
	return f[m]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(maxTaxiEarnings(5, [][]int{{2, 5, 4}, {1, 5, 1}}))
	//fmt.Println(maxTaxiEarnings(20, [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}))
	// 33
	fmt.Println(maxTaxiEarnings(10, [][]int{{2, 3, 6}, {8, 9, 8}, {5, 9, 7},
		{8, 9, 1}, {2, 9, 2}, {9, 10, 6}, {7, 10, 10}, {6, 7, 9}, {4, 9, 7}, {2, 3, 1}}))
}
