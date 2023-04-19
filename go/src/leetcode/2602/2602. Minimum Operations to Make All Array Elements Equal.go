package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/
*/
func minOperations(nums []int, queries []int) []int64 {
	ans := make([]int64, len(queries))
	sort.Ints(nums)

	n := len(nums)
	var sums = make([]int64, n+1)

	for i, x := range nums {
		sums[i+1] = sums[i] + int64(x)
	}

	for i, x := range queries {
		j := sort.SearchInts(nums, x) // nums[j] >= x
		ans[i] = int64(x*j) - sums[j] + (sums[n] - sums[j] - int64(x*(n-j)))

	}

	fmt.Println(ans)
	return ans
}

func main() {
	minOperations([]int{3, 1, 6, 8}, []int{1, 5})
	minOperations([]int{2, 9, 6, 3}, []int{10})
}
