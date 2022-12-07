/**
@author ZhengHao Lou
Date    2022/11/13
*/
package main

import "sort"

/**
https://leetcode.cn/problems/custom-sort-string/
*/
const N = 26

func customSortString(order string, s string) string {
	ord := make([]int, N)
	for i := range order {
		ord[int(order[i]-'a')] = i
	}
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return ord[int(bytes[i]-'a')] <= ord[int(bytes[j]-'a')]
	})
	return string(bytes)
}
