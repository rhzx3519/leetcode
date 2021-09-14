package main

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	var ans = 100001
	sort.Ints(nums)
	for i := 0; i + k - 1 < len(nums); i++ {
		ans = min(ans, nums[i+k-1] - nums[i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumDifference([]int{87063,61094,44530,21297,95857,93551,9918}, 6))
}