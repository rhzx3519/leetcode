package main

func maximumDifference(nums []int) int {
	n := len(nums)
	var diff = - 1

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] - nums[i] > 0 {
				diff = max(diff, nums[j] - nums[i])
			}
		}
	}

	return diff
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
