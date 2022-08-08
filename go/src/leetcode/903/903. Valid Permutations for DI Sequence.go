/**
@author ZhengHao Lou
Date    2022/02/09
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/valid-permutations-for-di-sequence/
思路：动态规划
f[i][j], 表示字符串s[:i]，排列结尾为j 时的有效排列个数
if s[i-1] == 'D', f[i][j] = sum(f[i-1][j ~ n])
if s[i-1] == 'I', f[i][j] = sum(f[i-1][0 ~ j])

f[0][0~n] = 1
*/
const MOD int = 1e9 + 7

func numPermsDISequence(s string) int {
	n := len(s)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := range f[0] {
		f[0][i] = 1
	}

	for i := 1; i <= n; i++ {
		if s[i-1] == 'D' { // p[i-1] > p[i]
			for j := 0; j <= i; j++ {
				for k := j; k < i; k++ {
					f[i][j] += f[i-1][k]
					f[i][j] %= MOD
				}
			}
		} else { // p[i-1] < p[i]
			for j := 0; j <= i; j++ {
				for k := 0; k < j; k++ {
					f[i][j] += f[i-1][k]
					f[i][j] %= MOD
				}
			}
		}
	}

	var c int
	for _, x := range f[n] {
		c += x
		c %= MOD
	}

	fmt.Println(c)
	return c
}

func main() {
	numPermsDISequence("DID")
}
