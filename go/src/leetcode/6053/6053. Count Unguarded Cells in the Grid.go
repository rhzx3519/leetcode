/**
@author ZhengHao Lou
Date    2022/05/01
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-unguarded-cells-in-the-grid/
*/

const (
	GUARD int = 1 << iota
	WALL
	RED
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for i := range guards {
		grid[guards[i][0]][guards[i][1]] = GUARD
	}
	for i := range walls {
		grid[walls[i][0]][walls[i][1]] = WALL
	}
	
	for i := 0; i < m; i++ {
		var g bool
		for j := 0; j < n; j++ {
			if grid[i][j] == GUARD {
				g = true
			} else if grid[i][j] == WALL {
				g = false
			} else if g {
				grid[i][j] |= RED
			}
		}
		g = false
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == GUARD {
				g = true
			} else if grid[i][j] == WALL {
				g = false
			} else if g {
				grid[i][j] |= RED
			}
		}
	}
	
	for j := 0; j < n; j++ {
		var g bool
		for i := 0; i < m; i++ {
			if grid[i][j] == GUARD {
				g = true
			} else if grid[i][j] == WALL {
				g = false
			} else if g {
				grid[i][j] |= RED
			}
		}
		g = false
		for i := m - 1; i >= 0; i-- {
			if grid[i][j] == GUARD {
				g = true
			} else if grid[i][j] == WALL {
				g = false
			} else if g {
				grid[i][j] |= RED
			}
		}
	}
	
	fmt.Println(grid)
	var c int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				c++
			}
		}
	}
	return c
}

func main() {
	countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}})
}
