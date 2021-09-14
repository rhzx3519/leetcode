package main

import "fmt"

/**
盈利方案数
假设你是一个交易员, 职责是买卖股票为你的组织带来盈利. 你每天只有一次操作的机会(买入或卖出), 并且在买入之前不能卖出.
现在给定一个数组, 代表某只股票每天的价格, 问随着价格变动, 一共有多少种盈利的方案?
难度
Medium
Example
input: [1, 2, 3]; output: 3
input: [1, 3, 2, 4]; output: 5
input: [6, 5, 4, 3]; output: 0
 */

func planNumber(prices []int) int {
	var count int

	var dfs func(idx, profit, lastOp int)
	dfs = func(idx, profit, lastOp int) {
		if idx == len(prices) {
			if lastOp == 1 && profit > 0 {
				count++
			}
			return
		}
		for i := idx; i < len(prices); i++ {
			op := lastOp*-1
			dfs(i+1, profit + op*prices[i], op)
		}
	}

	dfs(0, 0, 1)

	fmt.Println(count)
	return count
}

func main() {
	planNumber([]int{1,2,3})
}