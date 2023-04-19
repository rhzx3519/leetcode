package main

import (
	"fmt"
	"sort"
)

/**
思路：dp，dp[i]表示0...i最大的整除子集大小
dp[i] = max(dp[i], dp[j]+1), if nums[i] % nums[j] == 0
 */
func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	dp := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		parent[i] = -1
	}
	maxLen, maxIdx := -1, -1

	sort.Ints(nums)
	for i := 0; i < n; i++ {
		for j := i-1; j >= 0; j-- {
			if nums[i] % nums[j] != 0 {
				continue
			}
			if dp[j] + 1 > dp[i] {
				dp[i] = dp[j] + 1
				parent[i] = j
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
			maxIdx = i
		}
	}

	ans := []int{}
	for maxIdx != -1 {
		ans = append(ans, nums[maxIdx])
		maxIdx = parent[maxIdx]
	}
	return ans
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1,2,3}))
	fmt.Println(largestDivisibleSubset([]int{1,2,4,8}))
}
