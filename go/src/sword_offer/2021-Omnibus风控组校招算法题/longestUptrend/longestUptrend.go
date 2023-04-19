package main

import (
	"fmt"
	"math"
)

/**
上升通道
在交易中, 我们对拥有上涨趋势的股票称它 "处于上升通道中". 一般来讲, 该股票的股价只要不会跌破(<)上一个低点, 则有较大的概率继续上涨.
现在给定一个数组, 代表某只股票每天的价格. 问该股票处于上升通道的最大天数.
难度
Medium
Example
input: [1, 2, 3]; output: 3
input: [2, 1, 3, 1, 2, 5]; output: 5 ([1, 3, 1, 2, 5]因为每一次股价回落都没有跌破前一个低点, 股价由 3 跌到 1 的时候, 没有跌破 1(≤ 1))
input: [2, 3, 4, 1, 2, 1]; output: 3([2, 3, 4] or [1, 2, 1])

思路：
 */
func longestUptrend(prices []int) int {
	n := len(prices)
	var maxVal int
	for i := 0; i < n; i++ {
		var j = i
		for ; j < n && prices[j] >= prices[i]; j++ {
		}

		if j - i > maxVal {
			maxVal = j - i
		}
	}

	return maxVal
}

func longestUptrend2(prices []int) int {
	var s int
	minVal := math.MaxInt32
	for _, price := range prices {
		if price >= minVal {
			s++
		} else {
			s = 1
		}
		minVal = min(minVal, price)
	}
	return s
}

func longestUptrend3(prices []int) int {
	n := len(prices)
	var st = []int{}
	var dp = make([]int, n)
	for i := range dp {
		dp[i] = n
	}

	for i, price := range prices {
		for len(st) > 0 && prices[st[len(st)-1]] > price {
			dp[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	//fmt.Println(dp)
	var s int
	for i := range dp {
		s = max(s, dp[i] - i)
	}
	return s
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
	//fmt.Println(longestUptrend3([]int{1, 2, 3}))
	//fmt.Println(longestUptrend3([]int{2, 1, 3, 1, 2, 5}))
	fmt.Println(longestUptrend3([]int{2, 3, 4, 1, 2, 1}))
}
