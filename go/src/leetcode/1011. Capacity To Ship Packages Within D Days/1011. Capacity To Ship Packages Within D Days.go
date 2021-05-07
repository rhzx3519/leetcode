package main

import (
	"fmt"
	"math"
)

func shipWithinDays(weights []int, D int) int {
	l, r := max(weights), sum(weights)

	canShip := func(k, day int) bool {
		cur := k
		for _, w := range weights {
			if w > cur {
				cur = k
				day--
			}
			cur -= w
		}
		return day > 0
	}

	var mid int
	for l <= r {
		mid = l + (r-l)>>1
		if canShip(mid, D) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func sum(a []int) int {
	s := 0
	for _, x := range a {
		s += x
	}
	return s
}

func max(a []int) int {
	maxVal := math.MinInt32
	for _, x := range a {
		if x > maxVal {
			maxVal = x
		}
	}
	return maxVal
}

func main() {
	fmt.Println(shipWithinDays([]int{1,2,3,4,5,6,7,8,9,10}, 5))
	fmt.Println(shipWithinDays([]int{3,2,2,4,1,4}, 3))
	fmt.Println(shipWithinDays([]int{1,2,3,1,1}, 4))
}