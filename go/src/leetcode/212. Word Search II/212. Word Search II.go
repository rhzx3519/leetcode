package main

import "fmt"

var diff [][]int = [][]int{{1,0},{-1,0},{0,1},{0,-1}}

func findWords(board [][]byte, words []string) []string {
	ans := []string{}
	m, n := len(board), len(board[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}

	ws := make(map[string]int)
	for _, word := range words {
		ws[word]++
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp := []string{}
			for word, _ := range ws {
				if word[0] != board[i][j] {
					continue
				}
				visit[i][j] = true
				if appear(board, word, visit, i, j, 1) {
					visit[i][j] = false
					tmp = append(tmp, word)
				}
				visit[i][j] = false
			}
			for _, word := range tmp {
				for i := 0; i < ws[word]; i++ {
					ans = append(ans, word)
				}
				delete(ws, word)
			}
		}
	}
	return ans
}

func appear(board [][]byte, word string, visit [][]bool, x, y, idx int) bool {
	if idx == len(word)	{
		return true
	}

	m, n := len(board), len(board[0])
	for _, dxy := range diff {
		nx := x + dxy[0]
		ny := y + dxy[1]
		if nx >= 0 && nx < m && ny >= 0 && ny < n && !visit[nx][ny] && board[nx][ny] == word[idx] {
			visit[nx][ny] = true
			if appear(board, word, visit, nx, ny, idx+1) {
				visit[nx][ny] = false
				return true
			}
			visit[nx][ny] = false
		}
	}
	return false
}

func main() {
	fmt.Println(findWords([][]byte{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}},
		[]string{"oath","pea","eat","rain"}))
	fmt.Println(findWords([][]byte{{'a','b'},{'c','d'}}, []string{"abcb"}))
	fmt.Println(findWords([][]byte{{'a','a'}}, []string{"a"}))
}
