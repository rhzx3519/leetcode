package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/minimum-operations-to-reduce-x-to-zero/
思路：滑动窗口

 */
func minOperations(nums []int, x int) int {
	n := len(nums)
	var k = math.MaxInt32
	var s, l int

	for i := range nums {
		s += nums[i]
	}
	if s == x {
		return n
	}
	if s < x {
		return -1
	}
	x = s - x
	s = 0
	for r := 0; r < n; r++ {
		s += nums[r]
		for l <= r && s >= x {
			if s == x {
				k = min(k, n - (r - l + 1))
			}
			s -= nums[l]
			l++
		}
	}
	if k == math.MaxInt32 {
		return -1
	}
	return k
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minOperations([]int{1,1,4,2,3}, 5))
	fmt.Println(minOperations([]int{5,6,7,8,9}, 4))
	fmt.Println(minOperations([]int{3,2,20,1,1,3}, 10))
	fmt.Println(minOperations([]int{8828,9581,49,9818,9974,9869,9991,10000,10000,10000,9999,9993,9904,8819,1231,6309},
	134365))
}