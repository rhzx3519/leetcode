/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/buddy-strings/
 */
func buddyStrings(s string, goal string) bool {
	if s == goal {
		alphabet := make([]byte, 26)
		for i := 0; i < len(s); i++ {
			alphabet[int(s[i] - 'a')]++
			if alphabet[int(s[i] - 'a')] >= 2 {
				return true
			}
		}
	}
	if len(s) != len(goal) {
		return false
	}
	bytes := [][2]byte{}
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] != goal[i] {
			bytes = append(bytes, [2]byte{s[i], goal[i]})
		}
	}
	if len(bytes) != 2 {
		return false
	}
	//fmt.Println(bytes)
	return bytes[0][1] == bytes[1][0] && bytes[0][0] == bytes[1][1]
}

func main() {
	//fmt.Println(buddyStrings("ab", "ba"))
	//fmt.Println(buddyStrings("ab", "ab"))
	//fmt.Println(buddyStrings("aa", "aa"))
	//fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
	fmt.Println(buddyStrings("abcaa", "abcbb"))
}