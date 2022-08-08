/**
@author ZhengHao Lou
Date    2021/11/19
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/valid-tic-tac-toe-state/
 */
const N = 3

func validTicTacToe(board []string) bool {
	var wx, wo int

	var cx, co int
	for i := 0; i < N; i++ {
		var ro, rx int
		for j := 0; j < N; j++ {
			switch board[i][j] {
			case 'O':
				ro++
			case 'X':
				rx++
			}
		}
		if rx == N {
			wx++
		}
		if ro == N {
			wo++
		}
		co += ro
		cx += rx
	}

	if !check(cx, co, wx, wo) {
		return false
	}

	wx, wo = 0, 0
	cx, co = 0, 0
	for j := 0; j < N; j++ {
		var ro, rx int
		for i := 0; i < N; i++ {
			switch board[i][j] {
			case 'O':
				ro++
			case 'X':
				rx++
			}
		}
		if rx == N {
			wx++
		}
		if ro == N {
			wo++
		}
		co += ro
		cx += rx
	}
	if !check(cx, co, wx, wo) {
		return false
	}

	wx, wo = 0, 0
	var ro, rx int
	for i, j := 0, 0; i < N && j < N; i, j = i+1, j+1 {
		switch board[i][j] {
		case 'O':
			ro++
		case 'X':
			rx++
		}
	}
	if rx == N {
		wx++
	}
	if ro == N {
		wo++
	}
	if !check(cx, co, wx, wo) {
		return false
	}

	wx, wo = 0, 0
	ro, rx = 0, 0
	for i, j := 0, N-1; i < N && j >= 0; i, j = i+1, j-1 {
		switch board[i][j] {
		case 'O':
			ro++
		case 'X':
			rx++
		}
	}
	if rx == N {
		wx++
	}
	if ro == N {
		wo++
	}
	if !check(cx, co, wx, wo) {
		return false
	}

	return true
}

func check(cx, co, wx, wo int) bool {
	if cx != co && cx != co+1 {
		return false
	}
	if wo > 0 && wx > 0 {
		return false
	}
	if wx > 0 && cx == co {
		return false
	}
	if wo > 0 && cx == co+1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(validTicTacToe([]string{"O  ", "   ", "   "}))
	fmt.Println(validTicTacToe([]string{"XOX", " X ", "   "}))
	fmt.Println(validTicTacToe([]string{"XXX", "   ", "OOO"}))
	fmt.Println(validTicTacToe([]string{"XOX", "O O", "XOX"}))	// true
	fmt.Println(validTicTacToe([]string{"XO ","XO ","XO "}))	// false
	fmt.Println(validTicTacToe([]string{"XXX","OOX","OOX"}))	// true
	fmt.Println(validTicTacToe([]string{"XOX","OXO","XOX"}))	// true
	fmt.Println(validTicTacToe([]string{"XXX","XOO","OO "}))	// false
	fmt.Println(validTicTacToe([]string{"OXX","XOX","OXO"}))	// false
	fmt.Println(validTicTacToe([]string{"XXO","XOX","OXO"}))	// false
//
//["XXO",
// "XOX",
// "OXO"]
}