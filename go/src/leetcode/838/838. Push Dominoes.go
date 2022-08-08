/**
@author ZhengHao Lou
Date    2022/02/21
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/push-dominoes/
time: O(n)
space: O(n)
*/
func pushDominoes(dominoes string) string {
	n := len(dominoes)
	right := make([]int, n)

	var s int
	for i := range dominoes {
		c := dominoes[i]
		switch c {
		case 'R':
			s = 1
		case 'L':
			s = 0
		case '.':
			if s >= 1 {
				right[i] = s
				s++
			}
		}
	}

	bytes := []byte{}
	s = 0
	var b byte
	for i := n - 1; i >= 0; i-- {
		switch dominoes[i] {
		case 'R':
			b = 'R'
			s = 0
		case 'L':
			b = 'L'
			s = 1
		case '.':
			if s > 0 && (right[i] > s || right[i] == 0) {
				b = 'L'
			} else if right[i] == s {
				b = '.'
			} else {
				b = 'R'
			}
			if s > 0 {
				s++
			}
		}
		bytes = append(bytes, b)
	}
	reverse(bytes)

	fmt.Println(right)
	fmt.Println(string(bytes))
	return ""
}

func reverse(bytes []byte) {
	l, r := 0, len(bytes)-1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
}

func main() {
	pushDominoes("..")
	pushDominoes(".L.R...LR..L..")
}
