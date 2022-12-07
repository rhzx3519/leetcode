/**
@author ZhengHao Lou
Date    2022/10/14
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/distinct-subsequences-ii/
*/
const (
	N       = 26
	MOD int = 1e9 + 7
)

func distinctSubseqII(s string) (c int) {
	alphabeta := make([]int, N)
	for i := range s {
		var tot int
		for i := range alphabeta {
			tot = (tot + alphabeta[i]) % MOD
		}
		alphabeta[int(s[i]-'a')] += tot + 1
	}
	for i := range alphabeta {
		c = (c + alphabeta[i]) % MOD
	}
	return
}

func main() {
	fmt.Println(distinctSubseqII("abc"))
}
