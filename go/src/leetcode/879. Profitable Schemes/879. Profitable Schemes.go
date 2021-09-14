package main

import (
	"fmt"
	"math"
)

/**
dp[i][n]表示产生i的利润，并且工作的总数量最多为n的工作计划数量。
 */
var mod = int64(math.Pow10(9)) + int64(7)

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	var s int
	for _, p := range profit {
		s += p
	}
	dp := make([][]int, s+1)
	for i := 0; i <= s; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	m := len(group)
	for i := 0; i < m; i++ {
		for j := s; j >= profit[i]; j-- {
			for k := n; k >= group[i]; k-- {
				dp[j][k] = (dp[j][k] + dp[j-profit[i]][k-group[i]]) % int(mod)
			}
		}
	}

	var ans int64
	for i := minProfit; i <= s; i++ {
		for j := 0; j <= n; j++ {
			ans += int64(dp[i][j])
		}
	}
	return int(ans % mod)
}

func main() {
	fmt.Println(profitableSchemes(5, 3, []int{2,2}, []int{2,3}))
	fmt.Println(profitableSchemes(10, 5, []int{2,3,5}, []int{6,7,8}))
}
