/**
@author ZhengHao Lou
Date    2022/01/30
*/
package main

import "strings"

/**
https://leetcode-cn.com/problems/uncommon-words-from-two-sentences/
*/
func uncommonFromSentences(s1 string, s2 string) []string {
	m1, m2 := map[string]int{}, map[string]int{}
	for _, s := range strings.Split(s1, " ") {
		m1[s]++
	}
	for _, s := range strings.Split(s2, " ") {
		m2[s]++
	}
	var ans = []string{}
	for s, c := range m1 {
		if c == 1 && m2[s] == 0 {
			ans = append(ans, s)
		}
	}
	for s, c := range m2 {
		if c == 1 && m1[s] == 0 {
			ans = append(ans, s)
		}
	}
	return ans
}
