/**
@author ZhengHao Lou
Date    2022/10/30
*/
package main

/**
https://leetcode.cn/problems/words-within-two-edits-of-dictionary/
*/
func twoEditWords(queries []string, dictionary []string) (ans []string) {
OUT:
	for _, word := range queries {
		for _, target := range dictionary {
			var d int
			for i := 0; i < len(word); i++ {
				if word[i] != target[i] {
					d++
				}
			}
			if d <= 2 {
				ans = append(ans, word)
				continue OUT
			}
		}
	}
	return
}
