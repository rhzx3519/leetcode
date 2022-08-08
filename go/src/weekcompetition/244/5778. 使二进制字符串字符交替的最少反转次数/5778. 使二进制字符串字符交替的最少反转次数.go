package main

import "fmt"

/**
https://leetcode-cn.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/
思路：
1. 类型1操作就是将字符串s首位相连
2. 基于1，就是求s+s中，长度为len(s)的字符串变成交替字符串的最少类型2操作次数
3. 基于2，长度为len(s)的滑动窗口遍历s+s, 求窗口中字符串的最少类型2操作次数
4. 交替字符串可以是"1010", 也可以是"0101", 如果变成"1010"的操作次数是cnt，变成"0101"的操作次数为len(s) - cnt
5. 滑动窗口遍历时，如果出窗口的元素需要反转，则操作次数-1, 如果入窗口的元素需要反转，则操作次数+1
 */
func minFlips(s string) int {
	n := len(s)
	target := "01"
	var cnt int
	for i := range s {
		if s[i] != target[i%2] {
			cnt++
		}
	}
	var ans = min(cnt, n - cnt)
	s += s
	for i := 0; i < n; i++ {
		if s[i] != target[i%2] {
			cnt--
		}
		if s[i+n] != target[(i+n)%2] {
			cnt++
		}
		ans = min(min(ans, cnt), n - cnt)
	}

	fmt.Println(ans)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//minFlips("111000")
	minFlips("010")
	minFlips("1110")
	minFlips("01001001101")
}
