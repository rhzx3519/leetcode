package main

import (
	"fmt"
)

func jump(nums []int) int {
	n := len(nums)
	var step, maxDis, end int
	for i := 0; i < n-1; i++ {
		maxDis = max(maxDis, i+nums[i])
		if i == end {
			end = maxDis
			step++
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{2,3,0,1,4}))
}
