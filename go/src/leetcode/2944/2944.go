package main

import (
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/
思路: DP
f[i]表示一定要购买第i个元素所需要的最少金币数量
f[i] = price[i] + min(f[i+1, 2*i+1]), 其中i <= (n+2)/2
*/
func minimumCoins(prices []int) int {
	n := len(prices)
	f := make([]int, (n+1)/2)
	var dfs func(int) int
	dfs = func(i int) int {
		if i*2 >= n {
			return prices[i-1]
		}
		if f[i] != 0 {
			return f[i]
		}
		res := math.MaxInt32
		for j := i + 1; j <= i*2+1; j++ {
			res = min(res, dfs(j))
		}
		res += prices[i-1]
		f[i] = res
		return res
	}

	return dfs(1)
}

func dp(prices []int) int {
	n := len(prices)
	for i := (n+1)/2 - 1; i > 0; i-- {
		prices[i-1] += minArr(prices[i : i*2+1])
	}
	return prices[0]
}

func minArr(arr []int) int {
	var mi = math.MaxInt32
	for i := range arr {
		mi = min(mi, arr[i])
	}
	return mi
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumCoins([]int{3, 4, 1}))
}
