/**
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/stone-game-ii/
思路：动态规划
dp[i][j]表示剩余[i : len - 1]堆时，M = j的情况下，先取的人能获得的最多石子数

1. i + 2M >= len, dp[i][M] = sum[i : len - 1], 剩下的堆数能够直接全部取走，那么最优的情况就是剩下的石子总和
2. i + 2M < len, dp[i][M] = max(dp[i][M], sum[i : len - 1] - dp[i + x][max(M, x)]), 其中 1 <= x <= 2M，剩下的堆数不能全部取走，
那么最优情况就是让下一个人取的更少。对于我所有的x取值，下一个人从x开始取起，M变为max(M, x)，所以下一个人能取dp[i + x][max(M, x)]，
我最多能取sum[i : len - 1] - dp[i + x][max(M, x)]。

作者：lgh18
链接：https://leetcode-cn.com/problems/stone-game-ii/solution/java-dong-tai-gui-hua-qing-xi-yi-dong-17xing-by-lg/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func stoneGameII(piles []int) int {
	n := len(piles)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	var sum int
	for i := n-1; i >= 0; i-- {
		sum += piles[i]
		for M := 1; M <= n; M++ {
			if i + M<<1 >= n {
				f[i][M]	= sum
			} else {
				for x := 1; i+x < n && x <= M<<1; x++ {
					f[i][M] = max(f[i][M], sum - f[i+x][max(M, x)])
				}
			}
		}
	}
	return f[0][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(stoneGameII([]int{2,7,9,4,4}))
}