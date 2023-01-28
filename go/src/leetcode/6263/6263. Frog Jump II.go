/*
*
@author ZhengHao Lou
Date    2022/12/11
*/
package main

/*
*
https://leetcode.cn/problems/frog-jump-ii/
思路：贪心
采用间隔跳策略
*/
func maxJump(stones []int) int {
	var ans = stones[1] - stones[0]
	for i := 2; i < len(stones); i++ {
		ans = max(ans, stones[i]-stones[i-2])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
