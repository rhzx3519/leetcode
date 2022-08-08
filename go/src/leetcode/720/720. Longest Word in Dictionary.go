/**
@author ZhengHao Lou
Date    2022/03/17
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/longest-word-in-dictionary/
*/
func longestWord(words []string) string {
	mp := make(map[int][]string)
	for _, w := range words {
		k := len(w)
		if _, ok := mp[k]; !ok {
			mp[k] = []string{}
		}
		mp[k] = append(mp[k], w)
	}
	que := []string{}
	for _, w := range mp[1] {
		que = append(que, w)
	}

	var ans = ""
	var c int
	for len(que) != 0 {
		w := que[0]
		que = que[1:]
		if len(w) > c {
			c = len(w)
			ans = w
		} else if len(w) == c && w < ans {
			ans = w
		}

		k := len(w) + 1
		if _, ok := mp[k]; !ok {
			continue
		}
		for _, nw := range mp[k] {
			if w == nw[:k-1] {
				que = append(que, nw)
			}
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	//longestWord([]string{"w", "wo", "wor", "worl", "world"})
	longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})
}
