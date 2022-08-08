/**
@author ZhengHao Lou
Date    2021/11/15
*/
package main

/**
https://leetcode-cn.com/problems/check-whether-two-strings-are-almost-equivalent/
 */
const N = 26
func checkAlmostEquivalent(word1 string, word2 string) bool {
	c1 := statistic(word1)
	c2 := statistic(word2)
	for i := 0; i < N; i++ {
		if abs(c1[i] - c2[i]) > 3 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func statistic(word string) []int {
	c1 := make([]int, N)
	for _, c := range word {
		c1[int(c - 'a')]++
	}
	return c1
}
