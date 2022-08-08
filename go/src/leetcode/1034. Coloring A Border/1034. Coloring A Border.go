/**
@author ZhengHao Lou
Date    2021/12/07
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/coloring-a-border/
 */
var offset = [][]int{{1,0},{0,-1},{-1,0},{0,1}}
func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	c := grid[row][col]
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	que := []int{row*n + col}
	for len(que) != 0 {
		i := que[0]
		que = que[1:]
		x, y := i/n, i%n
		vis[x][y] = true
		var edge = x == m-1 || y == n-1 || x == 0 || y == 0
		for _, dxy := range offset {
			nx, ny := x + dxy[0], y + dxy[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || vis[nx][ny] {
				continue
			}
			if grid[nx][ny] != c {
				edge = true
			} else {
				que = append(que,  nx*n + ny)
			}
		}

		if edge {
			grid[x][y] = color
		}
	}

	fmt.Println(grid)
	return grid
}

func main() {
	colorBorder([][]int{{1,1},{1,2}}, 0, 0, 3)
	colorBorder([][]int{{1,2,2},{2,3,2}}, 0, 1, 3)
	colorBorder([][]int{{1,1,1},{1,1,1},{1,1,1}}, 1, 1, 2)
}
