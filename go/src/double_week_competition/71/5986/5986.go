/**
@author ZhengHao Lou
Date    2022/02/05
*/
package main

import "fmt"

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	var calc = func(min, sec int) int {
		var cost int
		digits := []int{min / 10, min % 10, sec / 10, sec % 10}
		var i int
		for ; digits[i] == 0; i++ {
		}
		digits = digits[i:]

		var last = startAt
		for _, d := range digits {
			cost += moveCost + pushCost
			if d == last {
				cost -= moveCost
			}
			last = d
		}
		return cost
	}

	min, sec := targetSeconds/60, targetSeconds%60
	if min >= 100 {
		min -= 1
		sec += 60
	}
	var cost = calc(min, sec)
	if min > 0 && sec+60 <= 99 {
		tmp := calc(min-1, sec+60)
		if tmp < cost {
			cost = tmp
		}
	}

	fmt.Println(cost)
	return cost
}

func main() {
	minCostSetTime(1, 2, 1, 600)
	minCostSetTime(0, 1, 2, 76)
	// 4
	minCostSetTime(9, 100000, 1, 6039)
	// 800000
	minCostSetTime(0, 100000, 100000, 5678)
}
