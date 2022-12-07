/*
*
@author ZhengHao Lou
Date    2022/11/27
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/count-palindromic-subsequences/
*/
const (
	MOD int = 1e9 + 7
	N       = 10
)

func countPalindromes(s string) (tot int) {
	n := len(s)
	var pre, suf = [N]int{}, [N]int{}
	var pre2, suf2 = [N][N]int{}, [N][N]int{}
	for i := n - 1; i >= 0; i-- {
		d := s[i] - '0'
		for j, c := range suf {
			suf2[d][j] += c
		}
		suf[d]++
	}
	for _, d := range s[:n-1] {
		d -= '0'
		suf[d]--
		for j, c := range suf {
			suf2[d][j] -= c
		}

		for j, sf := range suf2 {
			for k, c := range sf {
				tot += pre2[j][k] * c // jk d kj的组合数量
			}
		}
		for j, c := range pre {
			pre2[d][j] += c
		}
		pre[d]++
	}

	return tot % MOD
}

func main() {
	fmt.Println(countPalindromes("0000000"))
}
