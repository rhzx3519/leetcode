package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/
 */
func minDays(bloomDay []int, m, k int) int {
	if m > len(bloomDay)/k {
		return -1
	}
	minDay, maxDay := math.MaxInt32, 0
	for _, day := range bloomDay {
		if day < minDay {
			minDay = day
		}
		if day > maxDay {
			maxDay = day
		}
	}

	// 返回大于等于能够从花园中采摘m束花的最少天数
	l, r := minDay, maxDay
	for l < r {
		d := l + (r-l)>>1
		flowers, bouquet := 0, 0
		for _, day := range bloomDay {
			if day > d {
				flowers = 0
				continue
			}
			flowers++
			if flowers == k {
				flowers = 0
				bouquet++
			}
		}
		if bouquet >= m {
			r = d
		} else {
			l = d + 1
		}
	}

	return l
}

func lowerBound(arr []int, k int) int {
	l, r := 0, len(arr)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if arr[mid] >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(minDays([]int{1,10,3,10,2}, 3, 1))
	fmt.Println(minDays([]int{1,10,3,10,2}, 3, 2))
	fmt.Println(minDays([]int{7,7,7,7,12,7,7}, 2, 3))
}