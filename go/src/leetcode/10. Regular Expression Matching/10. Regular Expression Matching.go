package main

import "fmt"


func isMatch(s string, p string) bool {
	mem := make(map[[2]int]bool)
	var dp func(string, string, int, int) bool
	dp = func(s string, p string, i int, j int) bool {
		if b, ok := mem[[2]int{i, j}]; ok {
			return b
		}
		if j == len(p) {
			return i == len(s)
		}

		var b bool
		first := i < len(s) && (s[i] == p[j] || p[j] == '.')
		if j < len(p) - 1 && p[j+1] == '*' {
			// dp(i, j+2)表示p[j+1]='*'匹配0个s字符
			// first and dp(i+1, j)表示匹配1个或者多个字符
			b = dp(s, p, i, j + 2) || (first && dp(s, p, i + 1, j))
		} else {
			// 正常的相等匹配s[i]==p[j]
			b = first && dp(s, p, i + 1, j + 1)
		}
		mem[[2]int{i, j}] = b
		return b
	}

	return dp(s, p, 0, 0)
}



func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
