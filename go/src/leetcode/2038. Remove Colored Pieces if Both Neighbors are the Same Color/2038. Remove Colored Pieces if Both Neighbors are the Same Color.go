/**
@author ZhengHao Lou
Date    2021/10/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
 */
func winnerOfGame(colors string) bool {
	var count int
	var last byte
	var t = []int{0, 0}
	for i := range colors {
		c := colors[i]
		if c != last {
			if count >= 3 {
				t[int(last - 'A')] += count - 2
			}
			last = c
			count = 1
		} else {
			count++
		}
	}
	if count >= 3 {
		t[int(last - 'A')] += count - 2
	}
	fmt.Println(t)
	return t[0] > t[1]
}

func main() {
	winnerOfGame("AAABABB")
	winnerOfGame("AA")
	winnerOfGame("ABBBBBBBAAA")
	winnerOfGame("AAAAABBB")
}
