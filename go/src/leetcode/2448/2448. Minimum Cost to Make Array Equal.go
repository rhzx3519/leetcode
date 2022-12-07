/**
@author ZhengHao Lou
Date    2022/10/27
*/
package main

import "sort"

/**
https://leetcode.cn/problems/minimum-cost-to-make-array-equal/
思路：枚举
使所有数都等于nums[1],
1. 总开销tot 增加(nums[1] - nums[0]) * cost[0]
2. 总开销tot 减少(nums[1] - nums[0]) * (sumCost - cost[0])

总变化量tot += (nums[1] - nums[0) * (2*cost[0] - sumCost)
*/
func minCost(nums []int, cost []int) (ans int64) {
	n := len(nums)
	arr := make([][]int, n)
	for i := range nums {
		arr[i] = []int{nums[i], cost[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	var tot, sumCost int64
	for i := range arr {
		tot += int64((arr[i][0] - arr[0][0]) * arr[i][1])
		sumCost += int64(arr[i][1])
	}
	ans = tot
	for i := 1; i < n; i++ {
		sumCost -= int64(2 * arr[i-1][1])
		tot -= int64(arr[i][0]-arr[i-1][0]) * sumCost
		if tot < ans {
			ans = tot
		}
	}

	return ans
}
