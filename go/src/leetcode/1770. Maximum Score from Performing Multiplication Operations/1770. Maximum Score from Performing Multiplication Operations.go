package main

import (
	"fmt"
	"math"
)

var mem = make(map[[3]int]int)

func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	mem = make(map[[3]int]int)
	//return dp(0, 0, n-1,  nums, multipliers)
	dp := make([][]int, 1001)
	for i := 0; i < 1001; i++ {
		dp[i] = make([]int, 1001)
	}
	score := math.MinInt32
	for k := 1; k <= m; k++ {
		for i := 0; i <= k; i++ {
			if i == 0 {
				dp[i][k-i] = dp[i][k-i-1] + nums[n-k+i]*multipliers[k-1]
			} else if i == k {
				dp[i][k-i] = dp[i-1][k-i] + nums[i-1]*multipliers[k-1]
			} else {
				dp[i][k-i] = max(dp[i][k-i-1] + nums[n-k+i]*multipliers[k-1],
					dp[i-1][k-i] + nums[i-1]*multipliers[k-1])
			}

			if k == m {
				score = max(score, dp[i][k-i])
			}
		}
	}

	return score
}

func dp(idx, l, r int, nums, multipliers []int) int {
	if idx == len(multipliers) {
		return 0
	}
	key := [3]int{idx, l, r}
	if score, ok := mem[key]; ok {
		return score
	}

	mem[key] = max(dp(idx+1, l+1, r, nums, multipliers) + nums[l]*multipliers[idx],
		dp(idx+1, l, r-1, nums, multipliers) + nums[r]*multipliers[idx])
	return mem[key]
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maximumScore([]int{1,2,3}, []int{3,2,1}))
	fmt.Println(maximumScore([]int{-5,-3,-3,-2,7,1}, []int{-10,-5,3,4,6}))
}