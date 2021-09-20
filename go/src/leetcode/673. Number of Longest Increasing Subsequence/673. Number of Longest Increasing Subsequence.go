package main

import "fmt"

/**
思路：DP
f[i]表示以nums[i]结尾的最长递增子序列的长度
f[j] = max(f[j], f[i] + 1), nums[j] > nums[i], i < j
count[i]表示长度为f[i]的递增子序列的个数
 */
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	count := make([]int, n)
	for i := range nums {
		f[i] = 1
		count[i] = 1
	}

	var ans int
	maxVal := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if f[i] == f[j] + 1 {	//
					count[i] += count[j]
				} else if f[i] < f[j] + 1 {
					f[i] = f[j] + 1
					count[i] = count[j]
				}
			}
		}
		maxVal = max(maxVal, f[i])
	}

	fmt.Println(f, count, maxVal)
	for i := range f {
		if f[i] == maxVal {
			ans += count[i]
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findNumberOfLIS([]int{1,3,5,4,7}))
	fmt.Println(findNumberOfLIS([]int{2,2,2,2,2}))
}
