package main

/*
*
https://leetcode.cn/problems/longest-arithmetic-subsequence/
思路：动态规划
f[i][d]表示以nums[i]结尾的，等差为d的最长等差数列长度
*/
func longestArithSeqLength(nums []int) (tot int) {
	n := len(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 1001)
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j] + 500
			f[i][d] = max(f[i][d], f[j][d]+1)
			tot = max(f[i][d], tot)
		}
	}
	return tot + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
