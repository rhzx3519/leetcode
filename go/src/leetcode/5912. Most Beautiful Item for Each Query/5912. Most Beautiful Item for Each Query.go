/**
@author ZhengHao Lou
Date    2021/11/15
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/most-beautiful-item-for-each-query/
 */
func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	beauties := make([]int, len(items))
	var maxB int
	for i := range items {
		maxB = max(maxB, items[i][1])
		beauties[i] = maxB	// <=i's max beauty value
	}

	ans := []int{}
	for _, q := range queries {
		i := upperBound(items, q)
		if i == 0 {
			ans = append(ans, 0)
		} else {
			ans = append(ans, beauties[i-1])
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// return index where items[index][0] > x
func upperBound(items [][]int, x int) int {
	l, r := 0, len(items)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if items[m][0] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(maximumBeauty([][]int{{1,2},{3,2},{2,4},{5,6},{3,5}}, []int{1,2,3,4,5,6}))
}








