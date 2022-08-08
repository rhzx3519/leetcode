/**
@author ZhengHao Lou
Date    2022/06/23
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
*/
func findSubstring(s string, words []string) []int {
	var ans []int
	counter := make(map[string]int)
	for _, w := range words {
		counter[w]++
	}
	n := len(s)
	nw := len(words[0])
	nws := len(words)
	for i := 0; i+nws*nw <= n; i++ {
		t := copyMap(counter)
		for j := 0; j < nws; j++ {
			word := s[i+nw*j : i+(j+1)*nw]
			if t[word] == 0 {
				break
			}
			t[word]--
			if t[word] == 0 {
				delete(t, word)
			}
		}
		if len(t) == 0 {
			ans = append(ans, i)
		}
	}
	fmt.Println(ans)
	return ans
}

func copyMap(mp map[string]int) map[string]int {
	cp := make(map[string]int)
	for k, v := range mp {
		cp[k] = v
	}
	return cp
}

func main() {
	findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	findSubstring("barfoofoobarthefoobarman", []string{"foo", "bar", "the"})
	findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
	findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"})
}
