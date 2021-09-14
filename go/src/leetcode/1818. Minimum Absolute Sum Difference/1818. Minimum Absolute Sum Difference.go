package main

import (
	"fmt"
	"math"
	"sort"
)

/**
思路：abs(nums1[i] - nums2[i]) - abs(nums1[j] - nums2[i]) 求最大值
 */
var MOD = int(math.Pow(10, 9)) + 7

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	var maxDiff, sum int
	n := len(nums1)
	arr := append([]int{}, nums1...)
	sort.Ints(arr)
	var diff int
	for i := 0; i < n; i++ {
		x := abs(nums1[i] - nums2[i])
		sum += x
		j := lowerBound(arr, nums2[i])
		if j == n {
			diff = abs(arr[n-1] - nums2[i])
		} else if j == 0 {
			diff = abs(arr[0] - nums2[i])
		} else {
			diff = min(abs(arr[j] - nums2[i]), abs(arr[j-1] - nums2[i]))
		}
		maxDiff = max(maxDiff, x - diff)
	}
	return (sum - maxDiff) % MOD
}

func lowerBound(arr []int, x int) int {
	l, r := 0, len(arr)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if arr[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	fmt.Println(minAbsoluteSumDiff([]int{1,28,21}, []int{9,21,20}))
}
