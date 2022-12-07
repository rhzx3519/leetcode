/**
@author ZhengHao Lou
Date    2022/09/19
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/minimum-money-required-before-transactions/
*/
func minimumMoney(transactions [][]int) int64 {
	var ans, neg int64
	for i := range transactions {
		cost, cashback := int64(transactions[i][0]), int64(transactions[i][1])
		neg += max(0, cost-cashback)
	}
	for i := range transactions {
		cost, cashback := int64(transactions[i][0]), int64(transactions[i][1])
		ans = max(ans, neg+min(cost, cashback))
	}
	return ans
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumMoney([][]int{{2, 1}, {5, 0}, {4, 2}}))
	fmt.Println(minimumMoney([][]int{{3, 0}, {0, 3}}))
}
