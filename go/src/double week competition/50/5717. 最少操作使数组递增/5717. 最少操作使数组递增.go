package main

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	step := 0
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			step += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}

	return step
}

func main() {

	fmt.Println(minOperations([]int{1,1,1}))
	fmt.Println(minOperations([]int{1,5,2,4,1}))
	fmt.Println(minOperations([]int{8}))
}
