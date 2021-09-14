package main

import (
	"fmt"
	"math"
)

var mod = int(math.Pow10(9)) + 7

func kConcatenationMaxSum(arr []int, k int) int {
	s := sumArr(arr)
	m1 := maxArr(arr)%mod
	m2 := maxArr(append(arr, arr...))%mod
	if k == 1 {
		return m1
	} else if k == 2 {
		return m2
	}
	if s <= 0 {
		return max(m1, m2)
	}

	return max(m1, m2 + (k-2)*s%mod)%mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArr(arr []int) int {
	var maxVal = math.MinInt32
	s := 0
	for i := range arr {
		s += arr[i]
		if s < 0 {
			s = 0
		}
		if s > maxVal {
			maxVal = s
		}
	}
	return maxVal
}

func sumArr(arr []int) int {
	s := 0
	for i := range arr {
		s += arr[i]
	}
	return s
}

func main() {
	fmt.Println(kConcatenationMaxSum([]int{1,2}, 3))
	fmt.Println(kConcatenationMaxSum([]int{1,-2,1}, 5))
	fmt.Println(kConcatenationMaxSum([]int{-1,-2}, 7))
	fmt.Println(kConcatenationMaxSum([]int{1,1,-2}, 5))
	fmt.Println(kConcatenationMaxSum([]int{-5,4,-4,-3,5,-3}, 3))
	a := []int{-9,13,4,-16,-12,-16,3,-7,5,-16,16,8,-1,-13,15,3}
	fmt.Println(kConcatenationMaxSum(a, 6))
	fmt.Println(maxArr(a), maxArr(append(a, a...)), sumArr(a))
	fmt.Println(2000000000%mod)
}


