/**
@author ZhengHao Lou
Date    2022/05/23
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/minimum-lines-to-represent-a-line-chart/
*/
func minimumLines(stockPrices [][]int) (c int) {
	if len(stockPrices) < 2 {
		return 0
	}
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] <= stockPrices[j][0]
	})
	
	segs := []string{}
	
	for i := 0; i < len(stockPrices)-1; i++ {
		a, b := (stockPrices[i+1][1] - stockPrices[i][1]), (stockPrices[i+1][0] - stockPrices[i][0])
		g := gcd(a, b)
		segs = append(segs, fmt.Sprintf("%v/%v", a/g, b/g))
	}
	
	for i := 0; i < len(segs)-1; i++ {
		if segs[i] != segs[i+1] {
			c++
		}
	}
	c++
	return
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	//minimumLines([][]int{{1, 1000000000}, {1000000000, 1000000000}, {999999999, 1}, {2, 999999999}})
	//minimumLines([][]int{{1, 1}, {499999999, 2}, {999999998, 3}, {1000000000, 5}})
	minimumLines([][]int{{1, 1}, {500000000, 499999999}, {1000000000, 999999998}})
}
