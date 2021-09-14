package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-the-middle-index-in-array/
 */
func findMiddleIndex(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = nums[i-1] + sum[i-1]
	}
	var l, r int
	for i := 0; i < n; i++ {
		l = sum[i]
		r = sum[n] - sum[i+1]
		if l == r {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(findMiddleIndex([]int{2,3,-1,8,4}))
	fmt.Println(findMiddleIndex([]int{1,-1,4}))
	fmt.Println(findMiddleIndex([]int{2,5}))
	fmt.Println(findMiddleIndex([]int{1}))
}
