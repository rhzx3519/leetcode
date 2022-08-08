/**
@author ZhengHao Lou
Date    2022/05/10
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/cat-and-mouse-ii/
*/
const (
	K         = 1000
	S         = 8 * 8 * 8 * 8
	CAT_WIN   = 0
	MOUSE_WIN = 1
)

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type Pos struct {
	x, y int
}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	m, n := len(grid), len(grid[0])
	f := make([][]int, S)
	for i := range f {
		f[i] = make([]int, K)
		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var cx, cy, mx, my, fx, fy int
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 'C':
				cx, cy = i, j
			case 'M':
				mx, my = i, j
			case 'F':
				fx, fy = i, j
			}
		}
	}

	var dfs func(cx, cy, mx, my, k int) int
	dfs = func(cx, cy, mx, my, k int) int {
		state := (cx << 9) | (cy << 6) | (mx << 3) | my

		if k == K-1 {
			f[state][k] = CAT_WIN
			return f[state][k]
		}
		if cx == mx && cy == my {
			f[state][k] = CAT_WIN
			return f[state][k]
		}
		if mx == fx && my == fy {
			f[state][k] = MOUSE_WIN
			return f[state][k]
		}
		if cx == fx && cy == fy {
			f[state][k] = CAT_WIN
			return f[state][k]
		}

		if f[state][k] != -1 {
			return f[state][k]
		}

		if k&1 == 0 { // mouse
			for _, dir := range dirs {
				for d := 0; d <= mouseJump; d++ { // 它们可以停留在原地。
					nx := mx + dir[0]*d
					ny := my + dir[1]*d
					if nx < 0 || nx >= m || ny < 0 || ny >= n {
						break
					}
					if grid[nx][ny] == '#' { // 不能跳过墙
						break
					}
					if dfs(cx, cy, nx, ny, k+1) == MOUSE_WIN {
						f[state][k] = MOUSE_WIN
						return f[state][k]
					}
				}
			}
			f[state][k] = CAT_WIN
		} else { // mouse
			for _, dir := range dirs {
				for d := 0; d <= catJump; d++ { // 它们可以停留在原地。
					nx := cx + dir[0]*d
					ny := cy + dir[1]*d
					if nx < 0 || nx >= m || ny < 0 || ny >= n {
						break
					}
					if grid[nx][ny] == '#' {
						break
					}

					if dfs(nx, ny, mx, my, k+1) == CAT_WIN {
						f[state][k] = CAT_WIN
						return f[state][k]
					}
				}
			}
			f[state][k] = MOUSE_WIN
		}

		return f[state][k]
	}

	return dfs(cx, cy, mx, my, 0) == MOUSE_WIN
}

func main() {
	//fmt.Println(canMouseWin([]string{"####F", "#C...", "M...."}, 1, 2))
	//fmt.Println(canMouseWin([]string{"M.C...F"}, 1, 4))
	//fmt.Println(canMouseWin([]string{"....#..", "....#..", ".F.....", "#....M.", "...C...", "......#", "....###"},
	//	3, 3))
	fmt.Println(canMouseWin([]string{"F#M.", "..#.", "....", "..C.", "...#", "...."}, 3, 3))
}
