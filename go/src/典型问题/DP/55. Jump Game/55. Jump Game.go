package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	var cur int
	for i := 0; i < n; i++ {
		if cur >= i {
			cur = max(cur, i + nums[i])
		}
	}
	return cur >= n-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
}