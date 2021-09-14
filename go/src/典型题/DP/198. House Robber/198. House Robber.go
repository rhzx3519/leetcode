package main

import "fmt"

func rob(nums []int) int {
	var f1, f2 int
	var f3 int
	for i := range nums {
		f3 = max(f1 + nums[i], f2)
		f1 = f2
		f2 = f3
	}
	return f3
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1,2,3,1}))
	fmt.Println(rob([]int{2,7,9,3,1}))
}