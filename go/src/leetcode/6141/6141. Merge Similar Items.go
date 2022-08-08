/**
@author ZhengHao Lou
Date    2022/08/08
*/
package main

import "sort"

/**
https://leetcode.cn/problems/merge-similar-items/
*/
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	counter := make(map[int]int)
	for i := range items1 {
		counter[items1[i][0]] += items1[i][1]
	}
	for i := range items2 {
		counter[items2[i][0]] += items2[i][1]
	}
	var values []int
	for i := range counter {
		values = append(values, i)
	}

	sort.Ints(values)

	var ans [][]int
	for i := range values {
		ans = append(ans, []int{values[i], counter[values[i]]})
	}
	return ans
}
