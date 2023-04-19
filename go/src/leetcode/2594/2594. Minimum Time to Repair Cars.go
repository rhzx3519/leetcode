package main

import (
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/minimum-time-to-repair-cars/
r*n*n
*/
func repairCars(ranks []int, cars int) int64 {
	var mx int
	for _, r := range ranks {
		if r > mx {
			mx = r
		}
	}

	return Search(int64(mx)*int64(cars)*int64(cars), func(mintues int64) bool {
		n := cars
		for _, r := range ranks {
			n -= int(math.Sqrt(float64(mintues / int64(r))))
			if n <= 0 {
				break
			}
		}
		return n <= 0
	})
}

func Search(n int64, f func(int64) bool) int64 {
	var h, i, j int64
	i, j = 0, n
	for i < j {
		h = (i + j) >> 1
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func main() {
	fmt.Println(repairCars([]int{4, 2, 3, 1}, 10))
	fmt.Println(repairCars([]int{5, 1, 8}, 6))
	fmt.Println(repairCars([]int{3}, 52))
}
