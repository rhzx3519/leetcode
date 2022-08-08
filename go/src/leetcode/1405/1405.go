/**
@author ZhengHao Lou
Date    2022/02/07
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/longest-happy-string/
思路：贪心
time: O(a+b+c)
space: O(a+b+c)
*/
func longestDiverseString(a int, b int, c int) string {
	counter := map[byte]int{'a': a, 'b': b, 'c': c}
	words := []byte{'a', 'b', 'c'}

	var bytes []byte
	for flag := true; flag; {
		sort.Slice(words, func(i, j int) bool {
			return counter[words[i]] > counter[words[j]]
		})

		for i := range words {
			var w = words[i]
			if counter[w] <= 0 {
				flag = false
				break
			}

			if len(bytes) >= 2 && string(bytes[len(bytes)-2:]) == string([]byte{w, w}) {
				continue
			}

			bytes = append(bytes, w)
			counter[w] -= 1
			break
		}

	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	longestDiverseString(1, 1, 7)
	longestDiverseString(2, 2, 1)
	longestDiverseString(7, 1, 0)
}
