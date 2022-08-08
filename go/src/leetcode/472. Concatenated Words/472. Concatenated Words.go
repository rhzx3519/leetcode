/**
@author ZhengHao Lou
Date    2021/12/28
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/concatenated-words/
 */
func findAllConcatenatedWordsInADict(words []string) []string {
	counter := make(map[string]int)
	for _, w := range words {
		counter[w]++
	}

	var dp func(w string) bool
	dp = func(w string) bool {
		n := len(w)
		f := make([]int, n+1)	// f[i]表示下表从i开始的字符串可以连接的单词数量
		for i := range f {
			f[i] = -1
		}
		f[0] = 0
		var sub string
		for i := 0; i < n; i++ {
			if f[i] == -1 {
				continue
			}
			for j := i+1; j <= n; j++ {
				sub = w[i:j]
				if counter[sub] > 0 {
					f[j] = max(f[j], f[i] + 1)
				}
			}
			if f[n] > 1 {
				return true
			}
		}
		return false
	}

	var ans = []string{}
	for _, w := range words {
		counter[w]--
		if dp(w) {
			ans = append(ans, w)
		}
		counter[w]++
	}

	fmt.Println(ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	findAllConcatenatedWordsInADict([]string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"})
	findAllConcatenatedWordsInADict([]string{"cat","dog","catdog"})
	findAllConcatenatedWordsInADict([]string{""})
	findAllConcatenatedWordsInADict([]string{"a","b","ab","abc"})
}