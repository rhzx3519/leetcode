package main

import "fmt"

/**
https://leetcode-cn.com/problems/check-if-move-is-legal/
 */
const N = 8
var mapper = map[byte]int{'B': -1, 'W': 1, '.': 0}

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	for _, dxy := range [][]int{{-1,0},{1,0},{0,1},{0,-1},{-1,-1},{-1,1},{1,1},{1,-1}} {
		if checkLine(board, rMove, cMove, color, dxy) {
			return true
		}
	}
	return false
}

func checkLine(board [][]byte, r, c int, color byte, offset []int) bool {
	var l int
	i, j := r+offset[0], c+offset[1]
	for ; i >= 0 && i < N && j >= 0 && j < N; i, j = i+offset[0], j+offset[1] {
		if board[i][j] == '.' {
			break
		}
		if mapper[board[i][j]] ^ mapper[color] == 0 {
			break
		}
		l++
	}
	if l == 0 {
		return false
	}

	if i >= 0 && i < N && j >= 0 && j < N {
		return mapper[board[i][j]] ^ mapper[color] == 0
	}
	return false
}

func main() {
	//fmt.Println(checkMove([][]byte{{'.','.','.','B','.','.','.','.'},{'.','.','.','W','.','.','.','.'},{'.','.','.','W','.','.','.','.'},{'.','.','.','W','.','.','.','.'},{'W','B','B','.','W','W','W','B'},{'.','.','.','B','.','.','.','.'},{'.','.','.','B','.','.','.','.'},{'.','.','.','W','.','.','.','.'}},
	//	4, 3, 'B'))

	fmt.Println(checkMove([][]byte{{'.','.','.','.','.','.','.','.'},{'.','B','.','.','W','.','.','.'},{'.','.','W','.','.','.','.','.'},{'.','.','.','W','B','.','.','.'},{'.','.','.','.','.','.','.','.'},{'.','.','.','.','B','W','.','.'},{'.','.','.','.','.','.','W','.'},{'.','.','.','.','.','.','.','B'}},
	4, 4 ,'W'))
}
