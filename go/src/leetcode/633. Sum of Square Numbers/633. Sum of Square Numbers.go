package main

import "math"

func judgeSquareSum(c int) bool {
	l := 0
	r := int(math.Sqrt(float64(c)))
	var d int
	for l <= r {
		d = l*l + r*r
		switch {
		case d == c:
			return true
		case d > c:
			r--
		case d < c:
			l++
		}
	}
	return false
}
