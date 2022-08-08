/**
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/maximum-subarray-sum-with-one-deletion/
思路: dp
dp[i][0]表示不删除时，以i结尾的最大子数组和
dp[i][1]表示删除一个数时，以 i结尾的最大子数组和(分为删除arr[i] 和 删除i之前的一个数)
dp[i][0] = max(dp[i-1][0] + arr[i-1], arr[i-1])
dp[i][1] =  max(dp[i-1][1] + arr[i-1], dp[i-1][0])
time: O(n)
space: O(n)
 */

const NINF int = math.MinInt32 >> 1
func maximumSum(arr []int) int {
	var n = len(arr)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = []int{NINF, NINF}
	}
	var ans = NINF
	for i := 1; i <= n; i++ {
		f[i][0] = max(f[i-1][0] + arr[i-1], arr[i-1])
		f[i][1] = max(f[i-1][1] + arr[i-1], f[i-1][0])
		ans = max(ans, max(f[i][0], f[i][1]))
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	fmt.Println(maximumSum([]int{1,-2,0,3}))
	fmt.Println(maximumSum([]int{1,-2,-2,3}))
	fmt.Println(maximumSum([]int{-1,-1,-1,-1}))
}
