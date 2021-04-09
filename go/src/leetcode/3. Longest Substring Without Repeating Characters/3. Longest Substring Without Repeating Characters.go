package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	mem := make(map[int32]int)
	ans := 0
	last := 0
	for i, ch := range s {
		if _, ok := mem[ch]; ok {
			if mem[ch] + 1 > last {
				last = mem[ch] + 1
			}
		}
		//fmt.Println(i, last)

		if i - last + 1 > ans {
			ans = i - last + 1
		}
		mem[ch] = i
	}
	return ans
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
