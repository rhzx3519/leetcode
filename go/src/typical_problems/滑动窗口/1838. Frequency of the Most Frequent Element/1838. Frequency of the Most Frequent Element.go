package main

import (
	"fmt"
	"sort"
)

/**
思路：滑动窗口
 */
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	var d int
	var maxVal int
	for l, r := 0, 0; r < n; r++ {
		k -= (nums[r] - nums[l] - d) * (r - l)
		for l < r && k < 0 {
			k += nums[r] - nums[l]
			l++
		}
		d = nums[r] - nums[l]

		maxVal = max(maxVal, r - l + 1)
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxFrequency([]int{1,2,4}, 5))
	fmt.Println(maxFrequency([]int{1,4,8,13}, 5))
	fmt.Println(maxFrequency([]int{3,9,6}, 2))
}