/**
@author ZhengHao Lou
Date    2022/01/08
*/
package main

import (
	"fmt"
)

func longestPalindrome(words []string) int {
	var ans int
	counter := make(map[string]int)
	for _, w := range words {
		counter[w]++
	}

	keys := []string{}
	for key := range counter {
		keys = append(keys, key)
	}

	for _, w := range keys {
		if counter[w] == 0 {
			continue
		}
		image := string([]byte{w[1], w[0]})
		if counter[image] == 0 {
			continue
		}
		if image != w {
			c := min(counter[w], counter[image])
			ans += c<<2
			counter[w] -= c
			counter[image] -= c
			if counter[w] == 0 {
				delete(counter, w)
			}
			if counter[image] == 0 {
				delete(counter, image)
			}
		} else if counter[w] > 1 {
			c := counter[w]>>1
			ans += c<<2
			counter[w] -= c<<1
			if counter[w] == 0 {
				delete(counter, w)
			}
		}
	}

	for w := range counter {
		if w[0] == w[1] {
			ans += 2
			break
		}
	}

	fmt.Println(ans)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	longestPalindrome([]string{"lc","cl","gg"})
	longestPalindrome([]string{"ab","ty","yt","lc","cl","ab"})
	longestPalindrome([]string{"cc","ll","xx"})
}