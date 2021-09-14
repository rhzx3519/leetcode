package main

import "math"

func nthUglyNumber(n int, a int, b int, c int) int {
	var i, j, k int
	for i < n {
		minVal := min(i*a, j*b, k*c)
		switch minVal {
		case i * a:
			i++
		case j * b:
			j++
		case k * c:
			k++
		}

	}
	return 0
}

func min(a ...int) int {
	minVal := math.MaxInt32
	for i := range a {
		if minVal < a[i] {
			a[i] = minVal
		}
	}
	return minVal
}
