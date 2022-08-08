/**
@author ZhengHao Lou
Date    2021/11/17
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/walking-robot-simulation/
 */

var (
	N = 4
	dir = [][]int{{0,1},{1,0},{0,-1},{-1,0}}	// N,E,S,W
)

func robotSim(commands []int, obstacles [][]int) int {
	ob := make(map[[2]int]int)
	for i := range obstacles {
		ob[[2]int{obstacles[i][0], obstacles[i][1]}]++
	}

	var ans int
	var d = 0
	var x, y int
	for _, command := range commands {
		switch command {
		case -1:
			d = (d + 1) % N
		case -2:
			d = (d - 1 + N) % N
		default:
			var nx, ny = x, y
			for i := 0; i < command; i++ {
				if ob[[2]int{x + dir[d][0]*(i+1),y + dir[d][1]*(i+1)}] != 0 {
					break
				}
				nx = x + dir[d][0]*(i+1)
				ny = y + dir[d][1]*(i+1)
			}
			x, y = nx, ny
			ans = max(ans, x*x + y*y)
		}

	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(robotSim([]int{4,-1,3}, [][]int{}))
	//fmt.Println(robotSim([]int{4,-1,4,-2,4}, [][]int{{2,4}}))
	fmt.Println(robotSim([]int{6,-1,-1,6}, [][]int{}))	// 36
}
