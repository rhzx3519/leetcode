/**
@author ZhengHao Lou
Date    2022/06/07
*/
package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	var mx int
	for i := range piles {
		if piles[i] > mx {
			mx = piles[i]
		}
	}
	l, r := 1, mx
	for l < r {
		m := l + (r-l)>>1
		var t int
		for i := range piles {
			var x = piles[i] / m
			if piles[i]%m != 0 {
				x++
			}
			t += x
		}
		if t <= h {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
