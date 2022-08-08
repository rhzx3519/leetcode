/**
@author ZhengHao Lou
Date    2021/11/05
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/change-minimum-characters-to-satisfy-one-of-three-conditions/
 */
const N int = 26

func minCharacters(a string, b string) int {
	na, nb := len(a), len(b)
	ma := make([]int, 26)
	mb := make([]int, 26)
	for _, c := range a {
		ma[c - 'a']++
	}
	for _, c := range b {
		mb[c - 'a']++
	}

	var maxCount int
	for i := 0; i < N; i++ {
		maxCount = max(maxCount, ma[i] + mb[i])
	}

	var ans = na + nb - maxCount
	var ta, tb int
	for i := N - 1; i >= 0; i-- {
		ta += ma[i]
		tb += mb[i]
		ans = min(ans, min(ta + nb - tb, tb + na - ta))
		// nb - tb 表示b中小于i的字母个数,
		// ta + nb - tb 表示把b中小于i的字母替换成大于i的字母所需的最小操作次数
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCharacters("aba", "caa"))
	fmt.Println(minCharacters("dabadd", "cda"))
}

