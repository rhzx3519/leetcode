package main

import (
	"math"
)

/**
题目：https://leetcode-cn.com/problems/minimum-total-space-wasted-with-k-resizing-operations/
题解：https://leetcode-cn.com/problems/minimum-total-space-wasted-with-k-resizing-operations/solution/k-ci-diao-zheng-shu-zu-da-xiao-lang-fei-wxg6y/
思路：DP
题目等价于:
	给定数组 \textit{nums}nums 以及整数 kk，需要把数组完整地分成 k+1k+1 段连续的子数组，
	每一段的权值是「这一段的最大值乘以这一段的长度再减去这一段的元素和」。需要最小化总权值。

记 f[i][j]f[i][j] 表示将 nums[0..i] 分成 j 段的最小总权值。在进行状态转移时，
我们可以枚举最后的一段，那么就有状态转移方程：

	f[i][j] = min(f[i0-1][j-1] + g[i0][i])

其中 g[i0][i] 表示 nums[i0..i] 这一段的最大值乘以这一段的长度再减去这一段的元素和」。
在进行动态规划之前，我们可以使用二重循环预处理出所有的 g 值。

 */
func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	// g[i0][i] 表示nums[i0:i]这个子数组的权值, 即nums[i0:i]子数组的最大值乘以(i-i0+1) - sum(nums[i0:i])
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		maxVal := math.MinInt32 // 子数组最大值
		sum := 0                // 子数组的和
		for j := i; j < n; j++ {
			maxVal = max(maxVal, nums[j])
			sum += nums[j]
			g[i][j] = maxVal*(j-i+1) - sum
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+2)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= k+1; j++ {
			for i0 := 0; i0 <= i; i0++ {
				var last int
				if i0 > 0 {
					last = dp[i0-1][j-1]
				}
				dp[i][j] = min(dp[i][j], last + g[i0][i])
			}
		}
	}

	//fmt.Println(g)
	//fmt.Println(dp[n-1][k+1])
	return dp[n-1][k+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minSpaceWastedKResizing([]int{10,20}, 0)
	minSpaceWastedKResizing([]int{10,20,30}, 1)
	minSpaceWastedKResizing([]int{10,20,15,30,20}, 2)
}



