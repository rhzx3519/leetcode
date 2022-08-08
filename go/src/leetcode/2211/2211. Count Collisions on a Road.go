/**
@author ZhengHao Lou
Date    2022/03/21
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-collisions-on-a-road/
思路：分类讨论
time: O(n)
space: O(1)
*/
func countCollisions(directions string) int {
	var c int
	var r, s int
	for i := range directions {
		switch directions[i] {
		case 'R':
			r++
		case 'L':
			if r > 0 {
				c += 2
				c += r - 1
				r = 0
				s = 1
			} else if s > 0 {
				c++
			}
		case 'S':
			if r > 0 {
				c += r
				r = 0
			}
			s = 1
		}
	}

	fmt.Println(c)
	return c
}

func main() {
	//countCollisions("RLRSLL")
	//countCollisions("LLRR")

	countCollisions("LSSSLLSSSSLRRSLLLSLSLRRLLLLLRSSRSRRSLLLSSS") // 24
}
