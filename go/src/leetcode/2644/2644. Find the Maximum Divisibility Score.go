package main

import (
	"fmt"
	"math"
)

func maxDivScore(nums []int, divisors []int) int {
	var mx, ans = 0, math.MaxInt32
	for j := range divisors {
		var c int
		for i := range nums {
			if nums[i]%divisors[j] == 0 {
				c++
			}
		}
		if c > mx || (c == mx && divisors[j] < ans) {
			ans = divisors[j]
			mx = c
			fmt.Println(ans, c)
		}
	}
	return ans
}

func main() {
	maxDivScore([]int{4, 7, 9, 3, 9}, []int{5, 2, 3})
}
