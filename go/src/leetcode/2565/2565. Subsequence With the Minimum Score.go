package main

import "fmt"

/*
*
https://leetcode.cn/problems/subsequence-with-the-minimum-score/
思路：前后缀分解
surf[i]表示最长后缀结束的下标
pref[i]表示最长前缀结束的下标
*/
func minimumScore(s string, t string) int {
	n, m := len(s), len(t)
	surf := make([]int, n+1)
	surf[n] = m
	for i, j := n-1, m-1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		surf[i] = j + 1
	}
	var ans = surf[0]
	if ans == 0 { // t已经是s的一个子序列
		return 0
	}
	for i, j := 0, 0; i < n; i++ {
		if s[i] == t[j] { // j不会等于m，因为surf[0] != 0表示t不是s的子序列
			j++
			ans = min(ans, surf[i+1]-j) // 删除t[j:surf[i+1]]
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

func main() {
	fmt.Println(minimumScore("bhajhfgajhfehiebghcjifhaghaihgdeidifhghiajeadcdbcaibccfajeiffdibcdhedahcjahgafbijihjhabefedjebhiaefi",
		"bhajhfgajhfehiebghcjifhaghaihgdeidifhghiajeadcdbcaibccfajeiffdibcdhedahcjahgafbijihjhabefedjebhiaefia"))
}
