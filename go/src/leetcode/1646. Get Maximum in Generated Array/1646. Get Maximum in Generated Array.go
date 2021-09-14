package main

import "fmt"

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var maxVal int
	nums := make([]int, n + 1)
	nums[1] = 1
	for i := 2; i <= n; i++ {
		if i&1 == 1 {
			nums[i] = nums[i>>1] + nums[i>>1 + 1]
		} else {
			nums[i] = nums[i>>1]
		}
		maxVal = max(maxVal, nums[i])
	}
	fmt.Println(nums)
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	getMaximumGenerated(6)
}
