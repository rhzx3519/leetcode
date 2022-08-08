/**
@author ZhengHao Lou
Date    2022/05/16
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/maximum-consecutive-floors-without-special-floors/
*/
func maxConsecutive(bottom int, top int, special []int) (ans int) {
	sort.Ints(special)
	i := lowerBound(special, bottom)
	j := upperBound(special, top)
	if i == len(special) || j == 0 {
		return top - bottom + 1
	}
	ans = max(ans, special[i]-bottom)

	for k := i + 1; k < j; k++ {
		ans = max(ans, special[k]-special[k-1]-1)
	}
	ans = max(ans, top-special[j-1])

	fmt.Println(ans)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(nums []int, x int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	maxConsecutive(2, 9, []int{4, 6})
	maxConsecutive(6, 9, []int{7, 6, 8})
}
