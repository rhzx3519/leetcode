/**
@author ZhengHao Lou
Date    2021/11/25
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/moving-stones-until-consecutive/
 */
func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}

	var d, e int
	fmt.Println(a, b, c)
	d = b-a-1 + c-b-1
	if a+2 == b || b+2 == c {
		e = 1
	} else {
		e = min(b-a-1, 1) + min(c-b-1, 1)
	}
	return []int{e, d}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numMovesStones(1,2,5))
	fmt.Println(numMovesStones(4,3,2))
	fmt.Println(numMovesStones(3,5,1))	// 1,2
}
