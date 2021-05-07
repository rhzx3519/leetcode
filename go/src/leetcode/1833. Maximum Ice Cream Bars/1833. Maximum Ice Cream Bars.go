package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	var cnt int
	sort.Ints(costs)
	for _, cost := range costs {
		if coins < cost {
			break
		}
		coins -= cost
		cnt++
	}
	return cnt
}

