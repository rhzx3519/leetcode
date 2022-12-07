/**
@author ZhengHao Lou
Date    2022/10/20
*/
package main

/**
https://leetcode.cn/problems/k-th-symbol-in-grammar/
*/
func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	pn := kthGrammar(n-1, (k+1)/2)
	if pn == 0 {
		return k & 1
	}
	return 1 ^ (k & 1)
}
