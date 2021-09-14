package main

import (
	"fmt"
	"strconv"
)

/**
https://leetcode-cn.com/problems/form-largest-integer-with-digits-that-add-up-to-target/
思路：完全背包问题

 */
func largestNumber(cost []int, target int) string {
	dp := make([]string, target+1)
	dp[0] = ""
	for i := 1; i <= target; i++ {
		dp[i] = "#"
	}

	for t, cost := range cost {
		for i := cost; i <= target; i++ {
			// 当dp[i-cost]为"#"时，表示没有构造(t+1) + dp[i-cost]的路径
			if dp[i-cost] == "#" {
				continue
			}
			// cost数组遍历时是从低到高添加数位，当前遍历到t时，添加的数位肯定比dp[i-cost]大，
			// 也就是说构造的数字肯定时高位数字大的,比如"98722"
			dp[i] = max(dp[i], strconv.Itoa(t+1) + dp[i-cost])
		}
	}
	fmt.Println(dp[target])
	if dp[target] == "#" {
		return "0"
	}
	return dp[target]
}

func max(a, b string) string {
	if a == b {
		return a
	}
	n1, n2 := len(a), len(b)
	if n1 > n2 {
		return a
	} else if n1 < n2 {
		return b
	}
	for i := 0; i < n1; i++ {
		if a[i] < b[i] {
			return b
		} else if a[i] > b[i] {
			return a
		}
	}
	return a
}

func main() {
	largestNumber([]int{4,3,2,5,6,7,2,5,5}, 9)
	largestNumber([]int{7,6,5,5,5,6,8,7,8}, 12)
	largestNumber([]int{2,4,6,2,4,6,4,4,4}, 5)
}
