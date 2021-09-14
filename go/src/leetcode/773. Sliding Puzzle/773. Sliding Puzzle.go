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

	bytes := []byte{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			bytes = append(bytes, byte(board[i][j] + '0'))
		}
	}
	start := string(bytes)
	visited[start] = true
	que := []string{start}
	var step int
	for len(que) > 0 {
		for sz := len(que); sz > 0; sz -- {
			s := que[0]
			que = que[1:]

			if s == target {
				return step
			}
			i := strings.Index(s, "0")
			for _, ni := range neighbors[i] {
				bytes = []byte(s)
				bytes[i], bytes[ni] = bytes[ni], bytes[i]
				tmp := string(bytes)
				if _, ok := visited[tmp]; ok {
					continue
				}
				visited[tmp] = true
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
