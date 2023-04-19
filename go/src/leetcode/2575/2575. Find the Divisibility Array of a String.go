package main

/*
*
https://leetcode.cn/problems/find-the-divisibility-array-of-a-string/description/
*/
func divisibilityArray(word string, m int) []int {
	n := len(word)
	var ans = make([]int, n)
	var x int
	for i := range word {
		x = (x*10 + int(word[i]-'0')) % m
		if x == 0 {
			ans[i] = 1
		}
	}
	return ans
}
