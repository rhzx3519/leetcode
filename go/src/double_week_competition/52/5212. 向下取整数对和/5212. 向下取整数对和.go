package main

import (
	"fmt"
	"math"
	"sort"
)

var mod = int(math.Pow10(9)) + 7

func sumOfFlooredPairs(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	s := n
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			s = (s + floor(nums[j], nums[i])) % mod
		}
	}
	return s
}

func floor(a, b int) int {
	return a / b
}

func main() {
	fmt.Println(sumOfFlooredPairs([]int{2,5,9}))
	fmt.Println(sumOfFlooredPairs([]int{7,7,7,7,7,7,7}))
}
