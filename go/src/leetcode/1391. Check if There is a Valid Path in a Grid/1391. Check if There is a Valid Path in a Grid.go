package main

import "fmt"

var d = 4 // 上下左右, 0123
var offset = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
var mapper = make(map[int][]map[int]bool)


type pair struct {
	i, j int
}

func hasValidPath(grid [][]int) bool {
	var m, n = len(grid), len(grid[0])

	var que = []pair{{0, 0}}
	for len(que) != 0 {
		var p = que[0]
		que = que[1:]
		var street = grid[p.i][p.j]
		if street == -1 {
			continue
		}
		grid[p.i][p.j] = -1

		for dir, dxy := range offset {
			ni, nj := p.i + dxy[0], p.j + dxy[1]
			if ni < 0 || ni >= m || nj < 0 || nj >= n || grid[ni][nj] == -1 {
				continue
			}
			ns := grid[ni][nj]
			if _, ok := mapper[street][dir][ns]; !ok {
				continue
			}
			que = append(que, pair{ni, nj})
		}
	}

	return grid[m-1][n-1] == -1
}

func init() {
	for k := 1; k <= 6; k++ {
		mapper[k] = make([]map[int]bool, d)
		for j := 0; j < d; j++ {
			mapper[k][j] = make(map[int]bool)
		}
	}


	mapper[1][2][1] = true
	mapper[1][3][1] = true
	mapper[1][2][4] = true
	mapper[1][2][6] = true
	mapper[1][3][3] = true
	mapper[1][3][5] = true

	mapper[2][0][2] = true
	mapper[2][1][2] = true
	mapper[2][0][3] = true
	mapper[2][0][4] = true
	mapper[2][1][5] = true
	mapper[2][1][6] = true

	mapper[3][2][1] = true
	mapper[3][2][4] = true
	mapper[3][2][6] = true
	mapper[3][1][2] = true
	mapper[3][1][5] = true
	mapper[3][1][6] = true

	mapper[4][3][1] = true
	mapper[4][3][3] = true
	mapper[4][3][5] = true
	mapper[4][1][2] = true
	mapper[4][1][5] = true
	mapper[4][1][6] = true

	mapper[5][0][2] = true
	mapper[5][0][3] = true
	mapper[5][0][4] = true
	mapper[5][2][1] = true
	mapper[5][2][4] = true
	mapper[5][2][6] = true

	mapper[6][0][2] = true
	mapper[6][0][3] = true
	mapper[6][0][4] = true
	mapper[6][3][1] = true
	mapper[6][3][3] = true
	mapper[6][3][5] = true
}

func main() {
	//fmt.Println(hasValidPath([][]int{{2,4,3}, {6,5,2}}))
	//fmt.Println(hasValidPath([][]int{{1,2,1}, {1,2,1}}))
	//fmt.Println(hasValidPath([][]int{{1,2,1}, {1,2,1}}))
	//fmt.Println(hasValidPath([][]int{{1,1,2}}))
	//fmt.Println(hasValidPath([][]int{{1,1,1,1,1,1,3}}))
	//fmt.Println(hasValidPath([][]int{{2},{2},{2},{2},{2},{2},{6}}))

	fmt.Println(hasValidPath([][]int{{4,3,3},{6,5,2}}))
}