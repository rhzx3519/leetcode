/**
@author ZhengHao Lou
Date    2021/10/27
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/remove-invalid-parentheses/
思路：dfs + 剪枝
 */
func removeInvalidParentheses(s string) []string {
	var l, r int
	for i := range s {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' {
			r++
		}
	}

	n := len(s)
	var maxLen int
	MAX := min(l, r)
	mapper := make(map[string]int)

	var dfs func(idx int, t int, bytes []byte)
	dfs = func(idx int, t int, bytes []byte) {
		if t < 0 || t > MAX {
			return
		}

		if idx == n {
			if t == 0 && len(bytes) >= maxLen {
				str := string(bytes)
				if len(str) > maxLen {
					mapper = map[string]int{}
					maxLen = len(bytes)
				}
				mapper[str]++
			}
			return
		}

		var c = s[idx]
		if c == '(' {
			dfs(idx + 1, t + 1, append(append([]byte{}, bytes...), c))
			dfs(idx + 1, t, append([]byte{}, bytes...))
		} else if c == ')' {
			dfs(idx + 1, t - 1, append(append([]byte{}, bytes...), c))
			dfs(idx + 1, t, append([]byte{}, bytes...))
		} else {
			dfs(idx + 1, t, append(append([]byte{}, bytes...), c))
		}
	}

	dfs(0, 0, []byte{})
	ans := []string{}
	for s, _ := range mapper {
		ans = append(ans, s)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
	fmt.Println(removeInvalidParentheses(")("))
}
