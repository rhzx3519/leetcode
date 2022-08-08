/**
@author ZhengHao Lou
Date    2022/05/09
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/check-if-there-is-a-valid-parentheses-string-path/
*/
func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	mem := make([][][100]bool, m)
	for i := range mem {
		mem[i] = make([][100]bool, n)
	}
	
	var state bool
	var dfs func(x, y, c int)
	dfs = func(x, y, c int) {
		if state {
			return
		}
		
		if grid[x][y] == '(' {
			c++
		} else {
			c--
		}
		
		if c < 0 || c > m-1-x+n-1-y {
			return
		}
		
		if mem[x][y][c] {
			return
		}
		
		if x == m-1 && y == n-1 {
			state = c == 0
			return
		}
		
		if x+1 < m {
			dfs(x+1, y, c)
		}
		if y+1 < n {
			dfs(x, y+1, c)
		}
		
		mem[x][y][c] = true
		
		return
	}
	
	dfs(0, 0, 0)
	return state
}

func main() {
	fmt.Println(hasValidPath([][]byte{{'(', '(', '('}, {')', '(', ')'}, {'(', '(', ')'}, {'(', '(', ')'}}))
	fmt.Println(hasValidPath([][]byte{{')', ')'}, {'(', '('}}))
}
