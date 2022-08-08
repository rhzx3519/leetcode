/**
@author ZhengHao Lou
Date    2021/11/26
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/robot-bounded-in-circle/
 */
const N = 4
var dirs = [][]int{{0,1},{1,0},{0,-1},{-1,0}}
func isRobotBounded(instructions string) bool {
	n := len(instructions)
	mapper := map[[3]int]bool{}
	var d int
	var cur = [3]int{0,0,n-1}
	mapper[cur] = true
	for k := 0; k < 4; k++ {
		for i, c := range instructions {
			switch byte(c) {
			case 'G':
				cur[0] += dirs[d][0]*1
				cur[1] += dirs[d][1]*1
			case 'L':
				d = (d - 1 + N) % N
			case 'R':
				d = (d + 1)	% N
			}
			cur[2] = i%n
			//fmt.Println(cur)
			if mapper[cur] {
				return true
			}
			mapper[cur] = true
		}
	}

	return false
}

func main() {
	fmt.Println(isRobotBounded("GGLLGG"))
	fmt.Println(isRobotBounded("GG"))
	fmt.Println(isRobotBounded("GL"))
}