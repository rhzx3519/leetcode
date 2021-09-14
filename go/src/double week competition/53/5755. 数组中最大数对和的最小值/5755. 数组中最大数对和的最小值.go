package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1

	var ans int
	for l < r {
		ans = max(ans, nums[l] + nums[r])
		l++
		r--
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
	fmt.Println(minPairSum([]int{3,5,2,3}))
	fmt.Println(minPairSum([]int{3,5,4,2,4,6}))
}