package main

import "fmt"

/**
https://leetcode-cn.com/problems/longest-valid-parentheses/
 */
func longestValidParentheses(s string) int {
	n := len(s)
	f := make([]bool, n)
	st := []int{}
	for i := range s {
		if s[i] == '(' {
			st = append(st, i)
		} else if len(st) > 0 {
			j := st[len(st) - 1]
			st = st[:len(st) - 1]
			f[j], f[i] = true, true		// 标记有效的括号
		}
	}
	var count, maxVal int
	for i := 0; i < n; i++ {
		if f[i] {		// 格式正确且连续
			count++
		} else {
			maxVal = max(maxVal, count)
			count = 0
		}
	}

	fmt.Println(f)
	return max(maxVal, count)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	longestValidParentheses("(()")
	longestValidParentheses(")()())")
}