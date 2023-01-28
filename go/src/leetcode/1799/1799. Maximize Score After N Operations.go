/*
*
@author ZhengHao Lou
Date    2022/12/22
*/
package main

import "math/bits"

/*
*
https://leetcode.cn/problems/maximize-score-after-n-operations/description/
思路：状态压缩 + 动态规划
f[k] 表示当前状态为k时，能够获取的最大分数
f[k] = max(f[k], f[k^(1<<i)^(1<<j)] + c>>1*g[i][j]
*/
func maxScore(nums []int) int {
	m := len(nums)
	g := [14][14]int{}
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			g[i][j] = gcd(nums[i], nums[j])
		}
	}
	f := make([]int, 1<<m)
	for k := 0; k < 1<<m; k++ {
		c := bits.OnesCount(uint(k))
		if c&1 == 0 {
			for i := 0; i < m; i++ {
				if k>>i&1 == 1 {
					for j := 0; j < m; j++ {
						if k>>j&1 == 1 {
							f[k] = max(f[k], f[k^(1<<i)^(1<<j)]+c>>1*g[i][j])
						}
					}
				}
			}
		}
	}
	return f[1<<m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
