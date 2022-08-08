package main

import "fmt"

var offset = [][]int{{1,0},{-1,0},{0,1},{0,-1}}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	var id = 2
	var mapper = make(map[int]int)

	m, n := len(grid1), len(grid1[0])

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		var cnt = 1
		grid2[x][y] = id
		for _, dxy := range offset {
			nx := x + dxy[0]
			ny := y + dxy[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid2[nx][ny] == 1 {
				cnt += dfs(nx, ny)
			}
		}
		return cnt
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				k := dfs(i, j)
				mapper[id] = k
				id++
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 1 && grid2[i][j] != 0 {
				mapper[grid2[i][j]]--
			}
		}
	}

	var count int
	for _, cnt := range mapper {
		if cnt == 0 {
			count++
		}
	}
	fmt.Println(id, mapper, count)
	return count
}

func main() {
	countSubIslands([][]int{{1,1,1,0,0},{0,1,1,1,1},{0,0,0,0,0},{1,0,0,0,0},{1,1,0,1,1}},
	[][]int{{1,1,1,0,0},{0,0,1,1,1},{0,1,0,0,0},{1,0,1,1,0},{0,1,0,1,0}})

	countSubIslands([][]int{{1,0,1,0,1},{1,1,1,1,1},{0,0,0,0,0},{1,1,1,1,1},{1,0,1,0,1}},
		[][]int{{0,0,0,0,0},{1,1,1,1,1},{0,1,0,1,0},{0,1,0,1,0},{1,0,0,0,1}})
}
