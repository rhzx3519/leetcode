package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] += dp[i-c]
		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1,2,5}))
	fmt.Println(change(3, []int{2}))
	fmt.Println(change(10, []int{10}))
}
