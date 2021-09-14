package main

import "fmt"

/**
https://leetcode-cn.com/problems/stone-game-vii/
思路：二维动态规划
f[i][j]表示当石子剩余为stones[i,j]时，Alice and Bob's max diff
 */

func stoneGameVII(stones []int) int {
	n := len(stones)
	var sum = make([]int, n + 1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + stones[i-1]
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}

	for l := 2;  l <= n; l++ {
		for i := 0; i <= n - l; i++ {
			var j = i + l - 1
			left := sum[j+1] - sum[i+1]			// Alice got stones[i]
			right := sum[j] - sum[i]			// Alice got stones[j]
			f[i][j]= max(left - f[i+1][j], right - f[i][j-1])
		}
	}

	return f[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(stoneGameVII([]int{5,3,1,4,2}))
	fmt.Println(stoneGameVII([]int{7,90,5,1,100,10,10,2}))
}