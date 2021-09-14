package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	var f1, f2, f3 int
	for i := n-1; i >= 0; i-- {
		f3 = cost[i] + min(f1, f2)
		f1 = f2
		f2 = f3
	}
	return min(f1, f2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10,15,20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
