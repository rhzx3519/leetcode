/**
@author ZhengHao Lou
Date    2022/03/14
*/
package main

import "math"

/**
https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/
*/
func findRestaurant(list1 []string, list2 []string) []string {
	m1, m2 := map[string]int{}, map[string]int{}
	for i, s := range list1 {
		m1[s] = i
	}
	for i, s := range list2 {
		m2[s] = i
	}

	var ans []string
	var indexSum = math.MaxInt32
	for k, i := range m1 {
		if j, ok := m2[k]; ok {
			if i+j < indexSum {
				indexSum = i + j
				ans = []string{k}
			} else if i+j == indexSum {
				ans = append(ans, k)
			}
		}
	}

	return ans
}
