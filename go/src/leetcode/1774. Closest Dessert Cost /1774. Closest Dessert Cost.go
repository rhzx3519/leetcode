/*
思路：0-1背包问题，配药可以是0，1，2，所以扩展toppingCosts数组，重新添加一遍原来的元素
 */
package main

import (
	"fmt"
	"math"
)

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	N := 20000
	can := make([]bool, N+1)
	for _, base := range baseCosts {
		can[base] = true
	}

	toppingCosts = append(toppingCosts, toppingCosts...)
	for _, cost := range toppingCosts {
		for i := N; i >= cost; i-- {
			can[i] = can[i] || can[i-cost]
		}
	}

	diff := math.MaxInt32
	ans := 0
	for i := 1; i <= N; i++ {
		if can[i] && abs(i - target) < diff {
			diff = abs(i - target)
			ans = i
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(closestCost([]int{1,7}, []int{3,4}, 10))
	fmt.Println(closestCost([]int{2,3}, []int{4,5,100}, 18))
	fmt.Println(closestCost([]int{3,10}, []int{2,5}, 9))
	fmt.Println(closestCost([]int{10}, []int{1}, 1))
}


