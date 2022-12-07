/**
@author ZhengHao Lou
Date    2022/10/22
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/maximum-profit-in-job-scheduling/
思路：dp
dp[i] 表示兼职前i份工作的情况下能拿到的最大报酬
dp[i] = max(dp[i-1], profit[i-1] + dp[k])
*/
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = i
	}
	sort.Slice(jobs, func(i, j int) bool {
		if endTime[jobs[i]] != endTime[jobs[j]] {
			return endTime[jobs[i]] < endTime[jobs[j]]
		}
		if startTime[jobs[i]] != startTime[jobs[j]] {
			return startTime[jobs[i]] < startTime[jobs[j]]
		}
		return profit[jobs[i]] < profit[jobs[j]]
	})
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		k := sort.Search(n, func(j int) bool {
			return endTime[jobs[j]] > startTime[jobs[i-1]]
		})
		f[i] = max(f[i-1], profit[jobs[i-1]]+f[k])
	}
	fmt.Println(f)
	return f[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70})
	jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60})
	jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4})

	//var nums = []int{1, 3, 4, 6}
	//j := sort.Search(len(nums), func(i int) bool {
	//	return nums[i] > 4
	//})
	//fmt.Println(j)
}
