package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	st := []int{}
	for i := 0; i < 2*n; i++ {
		for len(st) > 0 && nums[st[len(st)-1]] < nums[i%n] {
			ans[st[len(st)-1]] = nums[i%n]
			st = st[:len(st)-1]
		}
		if i < n {
			st = append(st, i)
		}
	}

	return ans
}

func main() {
	fmt.Println(nextGreaterElements([]int{1,2,1}))
}