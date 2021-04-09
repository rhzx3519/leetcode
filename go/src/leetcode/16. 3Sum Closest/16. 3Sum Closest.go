package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	minVal := math.MaxInt32
	ans := 0
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		l, r := i+1, n-1
		for l < r {
			val := nums[i] + nums[l] + nums[r]
			if val < target {
				l++
			} else if val > target {
				r--
			} else {
				return target
			}
			if math.Abs(float64(val - target)) < float64(minVal) {
				minVal = int(math.Abs(float64(val - target)))
				ans = val
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(threeSumClosest([]int{-1,2,1,-4}, 1))
}