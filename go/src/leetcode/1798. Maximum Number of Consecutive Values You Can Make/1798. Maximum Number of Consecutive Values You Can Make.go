/**
思路：01背包问题
time: O(s*n), s为coins数组的和
space: O(s)
 */
package main

import (
	"fmt"
	"sort"
)

func getMaximumConsecutive(coins []int) int {
	n := len(coins)
	dp := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if i & (1<<j) != 0 {
				dp[i] += coins[j]
			}
		}
	}

	sort.Ints(dp)
	fmt.Println(dp)

	count := 0
	last := -1
	for _, num := range dp {
		if num == last+1 {
			last = num
			count++
		} else if num == last {

		} else {
			break
		}
	}

	return count
}

func main() {
	fmt.Println(getMaximumConsecutive([]int{1,3}))
	//fmt.Println(getMaximumConsecutive([]int{1,1,1,4}))
	//fmt.Println(getMaximumConsecutive([]int{1,4,10,3,1}))
}
