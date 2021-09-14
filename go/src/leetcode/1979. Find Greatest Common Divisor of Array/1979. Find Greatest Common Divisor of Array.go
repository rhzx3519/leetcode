package main

import "math"

func findGCD(nums []int) int {
	min, max := math.MaxInt32, math.MinInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return gcd(min, max)
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
