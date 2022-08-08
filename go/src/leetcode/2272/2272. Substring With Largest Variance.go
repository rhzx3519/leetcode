/**
@author ZhengHao Lou
Date    2022/05/16
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/substring-with-largest-variance/
思路：转化为子数组最大和https://leetcode.cn/problems/maximum-subarray/
假设出现次数最多的字符为a, 出现次数最少的字符为b，
可以认为s中字符a代表1，字符b代表-1,其余字符代表0，则题目转化为求子数组最大和
time: O(n)
space: O(1)
*/
func largestVariance(s string) (ans int) {
	n := len(s)
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}
			var diff, diffWithB = 0, -n
			for i := range s {
				if s[i] == byte(a) {
					diff++
					diffWithB++
				} else if s[i] == byte(b) {
					diff--
					diffWithB = diff
					diff = max(0, diff)
				}
				ans = max(ans, diffWithB)
			}
		}
	}

	fmt.Println(ans)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	largestVariance("aababbb")
	largestVariance("abcde")
}
