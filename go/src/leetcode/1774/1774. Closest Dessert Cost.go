/*
*
@author ZhengHao Lou
Date    2022/12/04
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/closest-dessert-cost/
*/
const N = 20000

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	can := make([]bool, N+1)
	for i := range baseCosts {
		can[baseCosts[i]] = true
	}
	toppingCosts = append(toppingCosts, toppingCosts...)
	for _, cost := range toppingCosts {
		for i := N; i >= cost; i-- {
			can[i] = can[i] || can[i-cost]
		}
	}
	var cost, d = -1, N
	for i := range can {
		if can[i] && abs(i-target) < d {
			d = abs(i - target)
			cost = i
		}
	}
	return cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(closestCost([]int{1, 7}, []int{3, 4}, 10))
	fmt.Println(closestCost([]int{2, 3}, []int{4, 5, 100}, 18))
	fmt.Println(closestCost([]int{3, 10}, []int{2, 5}, 9))
}
