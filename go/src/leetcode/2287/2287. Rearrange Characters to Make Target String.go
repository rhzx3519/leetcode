/**
@author ZhengHao Lou
Date    2022/05/30
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/rearrange-characters-to-make-target-string/
*/
const N = 26

func rearrangeCharacters(s string, target string) (c int) {
	c = len(s)
	letters := make([]int, N)
	for i := range s {
		letters[int(s[i]-'a')]++
	}
	t := make([]int, N)
	for i := range target {
		t[int(target[i]-'a')]++
	}
	
	fmt.Println(t)
	fmt.Println(letters)
	for i := range t {
		if t[i] == 0 {
			continue
		}
		if letters[i] == 0 {
			return 0
		}
		x := letters[i] / t[i]
		c = min(c, x)
	}
	
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	rearrangeCharacters("ilovecodingonleetcode", "code")
}
