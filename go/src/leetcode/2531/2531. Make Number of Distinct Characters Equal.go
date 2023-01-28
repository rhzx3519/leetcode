/**
@author ZhengHao Lou
Date    2023/01/09
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/make-number-of-distinct-characters-equal/
*/
func isItPossible(word1 string, word2 string) bool {
	var c1, c2 = make(map[rune]int), make(map[rune]int)
	for _, r := range word1 {
		c1[r]++
	}
	for _, r := range word2 {
		c2[r]++
	}
	
	for x, cx := range c1 {
		for y, cy := range c2 {
			// 交换x, y
			if x == y { // 如果不同字符串数目相同，此时交换x, y之后依然成立
				if len(c1) == len(c2) {
					return true
				}
			} else if len(c1)-b2i(cx == 1)+b2i(c1[y] == 0) == len(c2)-b2i(cy == 1)+b2i(c2[x] == 0) {
				return true
			}
		}
	}
	return false
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(isItPossible("abcde", "fghij"))
}
