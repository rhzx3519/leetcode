/**
@author ZhengHao Lou
Date    2022/05/04
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-the-winner-of-the-circular-game/
*/
func findTheWinner(n int, k int) int {
	friends := make([]int, n)
	for i := range friends {
		friends[i] = i + 1
	}
	var i = 0
	for len(friends) != 1 {
		i = (i + k - 1) % len(friends)
		fmt.Println(friends[i])
		copy(friends[i:], friends[i+1:])
		friends = friends[:len(friends)-1]
	}
	return friends[0]
}

func main() {
	findTheWinner(5, 2)
}
