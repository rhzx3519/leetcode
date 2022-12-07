/**
@author ZhengHao Lou
Date    2022/08/29
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/removing-stars-from-a-string/
*/
func removeStars(s string) string {
	var bytes []byte
	for i := range s {
		if s[i] != '*' {
			bytes = append(bytes, s[i])
		} else {
			if len(bytes) != 0 {
				bytes = bytes[:len(bytes)-1]
			}
		}
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	removeStars("leet**cod*e")
}
