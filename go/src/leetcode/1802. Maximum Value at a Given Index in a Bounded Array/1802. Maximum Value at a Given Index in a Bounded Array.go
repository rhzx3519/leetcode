package main

import (
	"fmt"
	"math"
)

func maxValue(n int, index int, maxSum int) int {
	ans := maxSum / n
	rest := maxSum % n
	l, r := index, index
	for l>0 || r<n-1 {
		length := r - l + 1
		if rest >= length {
			rest -= length
			l = max(0, l-1)
			r = min(n-1, r+1)
			ans++
		} else {
			break
		}

	}
	return ans
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func main() {
	fmt.Println(maxValue(4, 2, 6))
	fmt.Println(maxValue(6, 1, 10))
}