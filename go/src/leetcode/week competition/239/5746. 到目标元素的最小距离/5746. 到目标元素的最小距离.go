package main

import "math"

func getMinDistance(nums []int, target int, start int) int {
	diff := math.MaxInt32
	for i, num := range nums {
		if num == target {
			d := abs(i - start)
			if d < diff {
				diff = d
			}
		}
	}
	return diff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
