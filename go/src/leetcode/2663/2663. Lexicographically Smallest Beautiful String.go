package main

import "fmt"

/*
*
https://leetcode.cn/problems/lexicographically-smallest-beautiful-string/description/
*/
func smallestBeautifulString(S string, k int) string {
	limit := byte('a' + k)
	s := []byte(S)
	n := len(s)
	i := n - 1
	s[i]++

	for i < n {
		if s[i] == limit { // 超过范围
			if i == 0 { // 无法进位
				return ""
			}
			// 进位
			s[i] = 'a'
			i--
			s[i]++
		} else if i > 0 && s[i] == s[i-1] || i > 1 && s[i] == s[i-2] {
			s[i]++ // 如果 s[i] 和前面的字符形成回文串，就继续增加 s[i]
		} else {
			i++ // 检查 s[i] 是否和后面的字符形成回文串
		}
	}

	return string(s)
}

func main() {
	fmt.Println(smallestBeautifulString("abcz", 26))
	fmt.Println(smallestBeautifulString("dc", 4))
}
