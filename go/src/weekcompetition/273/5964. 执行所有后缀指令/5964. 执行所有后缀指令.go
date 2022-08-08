/**
@author ZhengHao Lou
Date    2021/12/26
*/
package main

import "fmt"

func executeInstructions(n int, startPos []int, s string) []int {
	ans := []int{}
	for i := 0; i < len(s); i++ {
		var x, y = startPos[0], startPos[1]
		var j = i
		for ; j < len(s); j++ {
			switch s[j] {
			case 'L':
				x, y = x, y-1
			case 'R':
				x, y = x, y+1
			case 'U':
				x, y = x-1, y
			case 'D':
				x, y = x+1, y
			}
			if x < 0 || x >= n || y < 0 || y >= n {
				break
			}
		}

		ans = append(ans, j - i)
	}
	fmt.Println(ans)
	return ans
}

func main() {
	executeInstructions(3, []int{0,1}, "RRDDLU")
}
