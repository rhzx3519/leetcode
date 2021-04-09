package main

import "fmt"

var mem = make(map[[2]int]bool)

func isMatch(s string, p string) bool {
	mem = make(map[[2]int]bool)
	return dp(s, p, 0, 0)
}

func dp(s, p string, i, j int) bool {
	if res, ok := mem[[2]int{i, j}]; ok {
		return res
	}
	if j == len(p) {
		return i == len(s)
	}
	ans := false
	first := i < len(s) && (p[j]==s[i] || p[j]=='.')
	if j < len(p)-1 && p[j+1] == '*' {
		// dp(i, j+2)表示p[j+1]='*'匹配0个s字符
		// first and dp(i+1, j)表示匹配1个或者多个字符
		ans = dp(s, p, i, j+2) || (first && dp(s, p, i+1, j))
	} else {
		// 正常的相等匹配s[i]==p[j]
		ans = first && dp(s, p, i+1, j+1)
	}
	mem[[2]int{i, j}] = ans
	return ans
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))

}
