package main

import "fmt"

/**
求一个数组子序列和的所有情况
https://leetcode-cn.com/problems/closest-subsequence-sum/
 */
func getSubSequenceSum(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	a := make([]int, 1 << (r + 1 - l))
	for i := l; i <= r; i++ {
		offset := 1 << (i - l)
		for j := 0; j < offset; j++ {
			a[j+offset] = a[j] + nums[i]
		}
	}
	return a
}

func main() {
	fmt.Println(getSubSequenceSum([]int{1,2,3}))
}

