package main

import (
	"fmt"
	"math"
)

var MOD = int(math.Pow(10, 9)) + 7

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	idx, diff := -1, 0
	n := len(nums2)
	s := 0
	for i := 0; i < n; i++ {
		d := abs(nums1[i] - nums2[i])
		s = (s + d) % MOD
		if d > diff {
			diff = d
			idx = i
		}
	}
	if idx == -1 {
		return 0
	}

	minDiff := diff
	var d int
	for i := 0; i < n; i++ {
		d = abs(nums1[i] - nums2[idx])
		if d < minDiff {
			minDiff = d
		}
	}

	return s - (diff - minDiff)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minAbsoluteSumDiff([]int{1,7,5}, []int{2,3,5}))
	fmt.Println(minAbsoluteSumDiff([]int{2,4,6,8,10}, []int{2,4,6,8,10}))
	fmt.Println(minAbsoluteSumDiff([]int{1,10,4,4,2,7}, []int{9,3,5,1,7,4}))
}
