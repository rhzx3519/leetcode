package main

import (
	"fmt"
	"sort"
)

func reductionOperations(nums []int) int {
	var n = len(nums)
	sort.Ints(nums)
	var s int
	var c = 1
	for i := n-1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			c++
		} else {
			s += n-i
			c = 1
		}
	}

	return s
}

func main() {
	fmt.Println(reductionOperations([]int{5,1,3}))
	fmt.Println(reductionOperations([]int{1,1,1}))
	fmt.Println(reductionOperations([]int{1,1,2,2,3}))
}
