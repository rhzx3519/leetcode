package main

import (
	"fmt"
	"strings"
)

var (
	neighbors map[int][]int
)

/**
	0 1 2
	3 4 5	-> 0 1 2 3 4 5

mapper return index i 's neighbors in two dimension array.
for instance, 4's neighbors in twon dimension is 1,3,5
 */
func init() {
	neighbors = make(map[int][]int)
	neighbors[0] = []int{1,3}
	neighbors[1] = []int{0,2,4}
	neighbors[2] = []int{1,5}
	neighbors[3] = []int{0,4}
	neighbors[4] = []int{1,3,5}
	neighbors[5] = []int{2,4}
}

func slidingPuzzle(board [][]int) int {
	var target = "123450"
	var visited = make(map[string]bool)

	var tmp = []rune{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			tmp = append(tmp, rune('0' + board[i][j]))
		}
	}
	var step int
	var start = string(tmp)
	var que = []string{start}
	for len(que) > 0 {
		sz := len(que)
		for ; sz > 0; sz-- {
			cur := que[0]
			que = que[1:]
			visited[cur] = true
			if cur == target {
				return step
			}
			i := strings.Index(cur, "0")
			var tmp string
			for _, nb := range neighbors[i] {
				tmp = cur[:i] + string(cur[nb]) + cur[i+1:]
				tmp = tmp[:nb] + string(cur[i]) + tmp[nb+1:]
				if _, ok := visited[tmp]; ok {
					continue
				}
				que = append(que, tmp)
			}
		}

		step++
	}

	return -1
}

func main() {
	fmt.Println(slidingPuzzle([][]int{{1,2,3}, {4,0,5}}))
	fmt.Println(slidingPuzzle([][]int{{1,2,3}, {5,4,0}}))
	fmt.Println(slidingPuzzle([][]int{{3,2,4}, {1,5,0}}))
}
