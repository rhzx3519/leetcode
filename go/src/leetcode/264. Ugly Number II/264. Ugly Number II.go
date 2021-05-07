package main

import "math"

func nthUglyNumber(n int) int {
	ugly := []int{1}
	var cur int
	var i, j, k int
	for count := 1;count < n; count++ {
		cur = min(ugly[i]*2, ugly[j]*3, ugly[k]*5)
		ugly = append(ugly, cur)
		for ;ugly[i]*2 <= cur; i++ {}
		for ;ugly[j]*3 <= cur; j++ {}
		for ;ugly[k]*5 <= cur; k++ {}
	}
	return ugly[len(ugly)-1]
}

func min(ints ...int) int {
	minVal := math.MaxInt32
	for _, i := range ints {
		if minVal > i {
			minVal = i
		}
	}
	return minVal
}

