/**
@author ZhengHao Lou
Date    2021/11/30
*/
package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters/
思路：贪心 + 单调栈
贪心策略：
1. 如果一个字符已经在栈中，则跳过
2. 如果栈顶的字符比当前字符大，且后面还有，则出栈
3. 最后栈中的就是最小的子序列
 */
const N = 26
func smallestSubsequence(s string) string {
	lastIndex := make([]int, N)
	for i := range lastIndex {
		lastIndex[i] = -1
	}
	for i := range s {
		lastIndex[int(s[i] - 'a')] = i
	}

	counter := make([]int, N)
	st := []byte{}
	for i := range s {
		c := s[i]
		if counter[int(c - 'a')] > 0 {
			continue
		}
		for len(st) > 0 && st[len(st) - 1] > s[i] && lastIndex[int(st[len(st) - 1] - 'a')] > i {
			counter[int(st[len(st) - 1] - 'a')]--
			st = st[:len(st) - 1]
		}
		st = append(st, c)
		counter[int(c - 'a')]++
	}

	fmt.Println(string(st))
	return string(st)
}

func main() {
	smallestSubsequence("bcabc")
	smallestSubsequence("cbacdcbc")
}
