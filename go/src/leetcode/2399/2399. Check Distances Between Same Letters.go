/**
@author ZhengHao Lou
Date    2022/09/05
*/
package main

/**
https://leetcode.cn/problems/check-distances-between-same-letters/
*/
func checkDistances(s string, distance []int) bool {
	counter := make(map[byte]int)
	for i := range s {
		if j, ok := counter[s[i]]; ok {
			if distance[int(s[i]-'a')] != i-j-1 {
				return false
			}
		}
		counter[s[i]] = i
	}
	return true
}
