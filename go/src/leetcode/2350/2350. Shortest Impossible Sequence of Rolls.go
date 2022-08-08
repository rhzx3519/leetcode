/**
@author ZhengHao Lou
Date    2022/07/25
*/
package main

/**
https://leetcode.cn/problems/shortest-impossible-sequence-of-rolls/
思路：贪心，遍历rolls每次集齐[1,k]所有数字时，表示序列长度为len的已经集齐了
*/
func shortestSequence(rolls []int, k int) int {
	var ans int
	var counter = map[int]int{}
	for i := range rolls {
		counter[rolls[i]]++
		if len(counter) == k {
			counter = map[int]int{}
			ans++
		}
	}
	return ans + 1
}
