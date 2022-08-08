/**
@author ZhengHao Lou
Date    2022/03/22
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
*/
func winnerOfGame(colors string) bool {
	var A, B int
	var a, b int
	for i := range colors {
		switch colors[i] {
		case 'A':
			a++
			if b >= 3 {
				B += b - 2
			}
			b = 0
		case 'B':
			b++
			if a >= 3 {
				A += a - 2
			}
			a = 0
		}
	}
	if a >= 3 {
		A += a - 2
	}
	if b >= 3 {
		B += b - 2
	}
	return A > B
}

func main() {
	fmt.Println(winnerOfGame("AAABABB"))
	fmt.Println(winnerOfGame("AA"))
	fmt.Println(winnerOfGame("ABBBBBBBAAA"))
}
