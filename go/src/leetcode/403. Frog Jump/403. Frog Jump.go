package main

import "fmt"

/*
   动态规划
   dp[i][k] 表示能否由前面的某一个石头 j 通过跳 k 步到达当前这个石头 i ，这个 j 的范围是 [1, i - 1]
   当然，这个 k 步是 i 石头 和 j 石头之间的距离
   那么对于 j 石头来说，跳到 j 石头的上一个石头的步数就必须是这个 k - 1 || k || k + 1
   由此可得状态转移方程：dp[i][k] = dp[j][k - 1] || dp[j][k] || dp[j][k + 1]
*/
func canCross(stones []int) bool {
	n := len(stones)
	dp := make([][]bool, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true;

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			k := stones[i] - stones[j]
			if k <= j + 1 {
				dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
				if i == n-1 && dp[i][k] {
					return true
				}
			}
		}
	}

	return false
}


func main() {
	fmt.Println(canCross([]int{0,1,3,5,6,8,12,17}))
}



