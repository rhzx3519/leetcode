/**
@author ZhengHao Lou
Date    2022/06/12
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/find-and-replace-pattern/
*/
func findAndReplacePattern(words []string, pattern string) []string {
	var ans []string
NEXT:
	for _, w := range words {
		mp := [2][26]byte{}
		n := len(w)
		for i := 0; i < n; i++ {
			ip, iw := index(pattern[i]), index(w[i])
			if mp[1][ip] != 0 && mp[0][iw] == 0 {
				continue NEXT
			}
			if mp[1][ip] == 0 && mp[0][iw] != 0 {
				continue NEXT
			}

			if (mp[1][ip] != 0 && mp[0][iw] != 0) && (mp[1][ip] != w[i] || mp[0][iw] != pattern[i]) {
				continue NEXT
			}
			mp[1][ip] = w[i]
			mp[0][iw] = pattern[i]
		}
		ans = append(ans, w)
	}
	return ans
}

func index(b byte) int {
	return int(b - 'a')
}

func main() {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}
