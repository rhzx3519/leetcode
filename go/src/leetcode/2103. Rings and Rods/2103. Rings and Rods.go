/**
@author ZhengHao Lou
Date    2021/12/13
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/rings-and-rods/
 */
func countPoints(rings string) int {
	rods := make(map[int]int)
	var c int
	for i := 0; i < len(rings); i += 2 {
		r := int(rings[i + 1] - '0')
		switch rings[i] {
		case 'R':
			rods[r] |= 1
		case 'G':
			rods[r] |= 1<<1
		case 'B':
			rods[r] |= 1<<2
		}
	}
	for _, color := range rods {
		if color == 7 {
			c++
		}
	}
	fmt.Println(rods, c)
	return c
}

func main() {
	//countPoints("B0B6G0R6R0R6G9")
	countPoints("B0R0G0R9R0B0G0")
	//countPoints("B0R0G0R9R0B0G0")
	//countPoints("G4")
}
