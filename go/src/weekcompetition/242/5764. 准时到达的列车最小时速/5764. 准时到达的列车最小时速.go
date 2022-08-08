package main

import (
	"fmt"
	"math"
)

/**
l = n
r = ceil(hour)
 */
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if hour < float64(n) - 1 {
		return -1
	}

	var mid int
	l, r := 1, int(math.Pow10(7))
	for l < r {
		mid = l + (r-l)>>1
		if check(dist, hour, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func check(dist []int, hour float64, speed int) bool {
	var s float64
	n := len(dist)
	if n == 1 {
		return float64(dist[0])/float64(speed) < hour
	}
	for i := 0; i < n-1; i++ {
		s += math.Ceil(float64(dist[i])/float64(speed))
	}
	s += float64(dist[n-1])/float64(speed)
	return float64(s) <= hour
}

func main() {
	fmt.Println(minSpeedOnTime([]int{1,3,2}, 6))
	fmt.Println(minSpeedOnTime([]int{1,3,2}, 2.7))
	fmt.Println(minSpeedOnTime([]int{1,3,2}, 1.9))
	fmt.Println(minSpeedOnTime([]int{1,1,100000}, 2.01))

}