/**
@author ZhengHao Lou
Date    2021/11/27
*/
package main

func countWords(words1 []string, words2 []string) int {
	var count int
	m1, m2 := statistic(words1), statistic(words2)
	for s1, c1 := range m1 {
		if c1 == 1 && m2[s1] == 1 {
			count++
		}
	}
	return count
}

func statistic(word []string) map[string]int {
	mp := make(map[string]int)
	for _, w := range word {
		mp[w]++
	}
	return mp
}
