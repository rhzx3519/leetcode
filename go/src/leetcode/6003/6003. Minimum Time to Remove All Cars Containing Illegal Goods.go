/**
@author ZhengHao Lou
Date    2022/02/07
*/
package main

/**
https://leetcode-cn.com/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/
思路：动态规划
pre[i]表示移除前i节的所有1车厢所花费的最小时间
if s[i] == '0' pre[i] = pre[i-1]
else pre[i] = min(pre[i-1] + 2, i + 1)

suf[i]表示移除后i节的所有1车厢所花费的最小时间
if suf[i] == '0', suf[i] =  suf[i+1]
else suf[i] = min(suf[i+1] + 2, n-i)
*/
func minimumTime(s string) int {
	n := len(s)
	var suf = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			suf[i] = suf[i+1]
		} else {
			suf[i] = min(suf[i+1]+2, n-i)
		}
	}
	var ans = suf[0]
	var pre int
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			pre = min(pre+2, i+1)
			ans = min(ans, pre+suf[i+1])
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
