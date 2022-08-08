/**
@author ZhengHao Lou
Date    2022/06/13
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/calculate-amount-paid-in-taxes/
*/
func calculateTax(brackets [][]int, income int) float64 {
	var tax float64
	for i, b := range brackets {
		var last = 0
		if i > 0 {
			last = brackets[i-1][0]
		}
		c := min(income, b[0]) - last
		tax += float64(c) * float64(b[1]) * 0.01
		if b[0] >= income {
			break
		}
	}
	fmt.Println(tax)
	return tax
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	calculateTax([][]int{{3, 50}, {7, 10}, {12, 25}}, 10)
}
