/**
@author ZhengHao Lou
Date    2022/11/10
*/
package main

import (
	"fmt"
	"unicode"
)

/**
https://leetcode.cn/problems/shortest-path-to-get-all-keys/
*/
var diff = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func shortestPathAllKeys(grid []string) int {
	var k int
	var sx, sy int
	m, n := len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if unicode.IsLower(rune(grid[i][j])) {
				k++
			} else if grid[i][j] == '@' {
				sx, sy = i, j
			}
		}
	}
	dist := make([][][]int, m)
	for i := range dist {
		dist[i] = make([][]int, n)
		for j := range dist[i] {
			dist[i][j] = make([]int, 1<<k)
			for t := range dist[i][j] {
				dist[i][j][t] = -1
			}
		}
	}

	dist[sx][sy][0] = 0
	var finished = 1<<k - 1
	var que = [][]int{{sx, sy, 0}}
	for len(que) != 0 {
		for tmp := len(que); tmp != 0; tmp-- {
			x, y, z := que[0][0], que[0][1], que[0][2]
			cur := dist[x][y][z]
			que = que[1:]
		NEXT:
			for _, d := range diff {
				nx, ny := x+d[0], y+d[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] == '#' {
					continue
				}

				if unicode.IsLower(rune(grid[nx][ny])) {
					i := int(grid[nx][ny] - 'a')
					nz := z | (1 << i)
					if nz == finished {
						return cur + 1
					}
					if dist[nx][ny][nz] != -1 {
						continue NEXT
					}
					dist[nx][ny][nz] = cur + 1
					que = append(que, []int{nx, ny, nz})
				} else if unicode.IsUpper(rune(grid[nx][ny])) {
					i := int(grid[nx][ny] - 'A')
					if (z & (1 << i)) == 0 {
						continue NEXT
					}
					if dist[nx][ny][z] != -1 {
						continue NEXT
					}
					dist[nx][ny][z] = cur + 1
					que = append(que, []int{nx, ny, z})
				} else {
					if dist[nx][ny][z] != -1 {
						continue NEXT
					}
					dist[nx][ny][z] = cur + 1
					que = append(que, []int{nx, ny, z})
				}
			}
		}
	}
	return -1
}

func main() {
	//fmt.Println(shortestPathAllKeys([]string{"@.a.#", "###.#", "b.A.B"}))
	//fmt.Println(shortestPathAllKeys([]string{"@..aA", "..B#.", "....b"}))
	//fmt.Println(shortestPathAllKeys([]string{"@Aa"}))
	//fmt.Println(shortestPathAllKeys([]string{"@abcdeABCDEFf"}))
	//fmt.Println(shortestPathAllKeys([]string{"@#....", ".Aabcd", ".BCD.."}))
	fmt.Println(shortestPathAllKeys([]string{"@...a", ".###A", "b.BCc"})) // 10
}
