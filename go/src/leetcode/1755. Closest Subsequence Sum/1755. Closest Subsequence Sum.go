package main

import (
	"fmt"
	"math"
	"sort"
)

func minAbsDifference(nums []int, goal int) int {
	ans := math.MaxInt32
	n := len(nums)

	var getSubSequenceSum func(l, r int, a []int)
	getSubSequenceSum = func(l, r int, a []int) {
		for i := l; i <= r; i++ {
			offset := 1 << (i - l)
			for j := 0; j < offset; j++ {
				a[j+offset] = a[j] + nums[i]
			}
		}
		sort.Ints(a)
	}

	var mid = n>>1
	var left = make([]int, 1<<(mid + 1 - 0))
	var right = make([]int, 1<<(n-1 + 1 - (mid+1)))
	getSubSequenceSum(0, mid, left)
	getSubSequenceSum(mid+1, n-1, right)

	sort.Ints(left)
	sort.Ints(right)

	l, r := 0, len(right)-1
	for l < len(left) && r >= 0 {
		val := left[l] + right[r]
		ans = min(ans, abs(goal - val))
		if val == goal {
			return 0
		} else if val > goal {
			r--
		} else {
			l++
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
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
	fmt.Println(minAbsDifference([]int{5,-7,3,5}, 6))
	fmt.Println(minAbsDifference([]int{7,-9,15,-2}, -5))
	fmt.Println(minAbsDifference([]int{1,2,3}, -7))
}
