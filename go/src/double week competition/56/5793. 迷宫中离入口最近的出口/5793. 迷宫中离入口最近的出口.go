package main

import "fmt"

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])

	que := [][]int{entrance}
	maze[entrance[0]][entrance[1]] = '+'
	var step int
	for len(que) > 0 {
		for sz := len(que); sz > 0; sz-- {
			e := que[0]
			que = que[1:]
			x, y := e[0], e[1]

			for _, dxy := range([][]int{{1,0},{-1,0},{0,1},{0,-1}})	 {
				nx := x + dxy[0]
				ny := y + dxy[1]
				if nx >= 0 && nx <= m-1 && ny >= 0 && ny <= n-1 && maze[nx][ny] == '.' {
					if nx == 0 || nx == m-1 || ny == 0 || ny == n-1 {
						return step+1
					}
					maze[x][y] = '+'
					que = append(que, []int{nx, ny})
				}
			}
		}
		step++
	}
	return -1
}

func main() {
	fmt.Println(nearestExit([][]byte{{'+','+','.','+'},{'.','.','.','+'},{'+','+','+','.'}}, []int{1,2}))
}
