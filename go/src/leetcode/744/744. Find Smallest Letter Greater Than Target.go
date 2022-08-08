/**
@author ZhengHao Lou
Date    2022/04/03
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	l, r := 0, n
	var m int
	for l < r {
		m = l + (r-l)>>1
		if letters[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l == n {
		return letters[0]
	}
	return letters[l]
}

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd')))
	fmt.Println(string(nextGreatestLetter([]byte{'a', 'b'}, 'z')))
}
