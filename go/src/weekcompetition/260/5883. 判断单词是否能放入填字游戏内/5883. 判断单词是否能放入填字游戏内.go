package main

import "fmt"

func placeWordInCrossword(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' || (board[i][j] != ' ' && board[i][j] != word[0]) {
				continue
			}
			// horizontal
			if j == 0 || isValid(board[i][j-1]) {
				var k = 0
				for k < min(n - j, len(word)) {
					if board[i][k + j] != ' ' && board[i][k + j] != word[k] {
						break
					}
					k++
				}
				if k == len(word) && (k + j == n || isValid(board[i][k+j])) {
					return true
				}
			}

			if j == n-1 || isValid(board[i][j+1]) {
				var k = 0
				for k < len(word) && j - k >= 0 {
					if board[i][j - k] != ' ' && board[i][j - k] != word[k] {
						break
					}
					k++
				}
				if k == len(word) && (j - k < 0 || isValid(board[i][j - k])) {
					return true
				}
			}

			// vertical
			if i == 0 || isValid(board[i-1][j]) {
				var k = 0
				for k < min(m - i, len(word)) {
					if board[k+i][j] != ' ' && board[k+i][j] != word[k] {
						break
					}
					k++
				}

				if k == len(word) && (k+i == m || isValid(board[k+i][j])) {
					return true
				}
			}
			if i == m-1 || isValid(board[i+1][j]) {
				var k = 0
				for k < len(word) && i - k >= 0 {
					if board[i-k][j] != ' ' && board[i-k][j] != word[k] {
						break
					}
					k++
				}

				if k == len(word) && (i - k < 0 || isValid(board[i-k][j])) {
					return true
				}
			}
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isValid(b byte) bool {
	if b == ' ' {
		return false
	}
	if b >= 'a' && b <= 'z' {
		return false
	}
	return true
}

func main() {
	//fmt.Println(placeWordInCrossword([][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', 'c', ' '}}, "abc"))
	fmt.Println(placeWordInCrossword([][]byte{{' ', '#', 'a'}, {' ', '#', 'c'}, {' ', '#', 'a'}}, "ac"))
	//fmt.Println(placeWordInCrossword([][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', ' ', 'c'}}, "ca"))
	//fmt.Println(placeWordInCrossword([][]byte{{' '},{'#'},{'o'},{' '},{'t'},{'m'},{'o'},{' '},{'#'},{' '}},
	//"octmor"))
	//fmt.Println(placeWordInCrossword([][]byte{{'#',' ','#'},{'#',' ','#'},{'#',' ','c'}}, "ca"))
}
