package main

import (
	"math"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		} else {
			return envelopes[i][1] > envelopes[j][1]
		}
	})
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = envelopes[i][1]
	}
	return LIS(nums)
}

func LIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = 1
	}

	maxVal := 0
	for i := 1; i < n+1; i++ {
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	maxEnvelopes([][]int{{5,4},{6,4},{6,7},{2,3}})
}

//思路：
//先对宽度 w 进行升序排序，如果遇到 w 相同的情况，则按照高度 h 降序排序。之后把所有的 h 作为一个数组，在这个数组上计算 LIS 的长度就是答案。
//这个解法的关键在于，对于宽度 w 相同的数对，要对其高度 h 进行降序排序。因为两个宽度相同的信封不能相互包含的，逆序排序保证在 w 相同的数对中最多只选取一个。
