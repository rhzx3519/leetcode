/**
@author ZhengHao Lou
Date    2022/01/22
*/
package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	var money int
	n := len(cost)
	sort.Ints(cost)
	var i int
	for i = n - 1; i-2 >= 0; i -= 3 {
		money += cost[i] + cost[i-1]
	}

	for ; i >= 0; i-- {
		money += cost[i]
	}

	return money
}

func main() {
	fmt.Println(minimumCost([]int{1}))
	fmt.Println(minimumCost([]int{6, 5, 7, 9, 2, 2}))
	fmt.Println(minimumCost([]int{5, 5}))
}
