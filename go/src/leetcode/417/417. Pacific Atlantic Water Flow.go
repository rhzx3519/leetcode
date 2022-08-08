/**
@author ZhengHao Lou
Date    2022/04/27
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/pacific-atlantic-water-flow/
*/

var d = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific, atlantic := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}
	bfs(heights, pacific, true)
	bfs(heights, atlantic, false)
	
	var ans [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	
	return ans
}

func bfs(heights [][]int, grid [][]bool, isPacific bool) {
	m, n := len(heights), len(heights[0])
	var que [][]int
	for i := 0; i < m; i++ {
		if isPacific {
			que = append(que, []int{i, 0})
			grid[i][0] = true
		} else {
			que = append(que, []int{i, n - 1})
			grid[i][n-1] = true
		}
	}
	
	for j := 0; j < n; j++ {
		if isPacific {
			que = append(que, []int{0, j})
			grid[0][j] = true
		} else {
			que = append(que, []int{m - 1, j})
			grid[m-1][j] = true
		}
	}
	
	for len(que) != 0 {
		g := que[0]
		que = que[1:]
		x, y := g[0], g[1]
		//fmt.Println(x, y)
		for _, nd := range d {
			nx, ny := nd[0]+x, nd[1]+y
			if nx >= 0 && nx < m && ny >= 0 && ny < n &&
				!grid[nx][ny] && heights[nx][ny] >= heights[x][y] {
				grid[nx][ny] = true
				que = append(que, []int{nx, ny})
			}
		}
	}
	fmt.Println()
}

func main() {
	pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}})
}
