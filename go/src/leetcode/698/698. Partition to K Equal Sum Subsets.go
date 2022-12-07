/**
@author ZhengHao Lou
Date    2022/09/20
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/
*/
func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	var tot int
	for i := range nums {
		tot += nums[i]
	}
	if tot%k != 0 {
		return false
	}
	var g = tot / k
	var vis = make([]bool, n)
	var dfs func(idx, s, c int) bool
	dfs = func(idx, s, c int) bool {
		if c == 0 {
			return true
		}
		if s == 0 {
			return dfs(0, g, c-1)
		}

		for i := idx; i < n; i++ {
			if vis[i] {
				continue
			}

			vis[i] = true
			if s >= nums[i] && dfs(i+1, s-nums[i], c) {
				return true
			}
			vis[i] = false
		}
		return false
	}

	return dfs(0, g, k)
}

func main() {
	//fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	//fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
	fmt.Println(canPartitionKSubsets([]int{129, 17, 74, 57, 1421, 99, 92, 285, 1276, 218, 1588, 215, 369, 117, 153, 22}, 3))
}
