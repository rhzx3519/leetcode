package main

/*
*
https://leetcode.cn/problems/extra-characters-in-a-string/?envType=daily-question&envId=2024-01-09
思路：动态规划
*/
func minExtraChar(s string, dictionary []string) int {
	memo := make(map[string]bool)
	for _, d := range dictionary {
		memo[d] = true
	}
	n := len(s)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] + 1
		for j := 0; j < i; j++ {
			if memo[s[j:i]] && f[j] < f[i] {
				f[i] = f[j]
			}
		}
	}
	return f[n]
}
