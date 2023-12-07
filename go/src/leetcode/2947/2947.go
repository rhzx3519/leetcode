package main

import "fmt"

/*
*
https://leetcode.cn/problems/count-beautiful-substrings-i/description/
*/
var vowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func beautifulSubstrings(s string, k int) (tot int) {
	var x int
	for i := range s {
		if vowels[s[i]] {
			x++
		}
		var y int
		for j := 0; j < i; j++ {
			if i-j+1 == (x-y)*2 && (x-y)*(x-y)%k == 0 {
				tot++
			}
			if vowels[s[j]] {
				y++
			}
		}
	}

	return
}

func main() {
	fmt.Println(beautifulSubstrings("baeyh", 2))
	fmt.Println(beautifulSubstrings("abba", 1))
	fmt.Println(beautifulSubstrings("bcdf", 1))
}
