package main

import "fmt"

/**
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 */
func lengthOfLongestSubstring(s string) int {
	var l, last int
	mapper := make(map[byte]int)
	for i := range s {
		c := s[i]
		if j, ok := mapper[c]; ok {
			last = max(j + 1, last)
		}
		l = max(l, i - last + 1)
		mapper[c] = i
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main()  {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("abc"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
