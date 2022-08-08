/**
@author ZhengHao Lou
Date    2021/12/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/battleships-in-a-board/
思路：从左上往右下遍历，统计战舰左上角的个数，
 */
func countBattleships(board [][]byte) int {
	var c int
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' && (i==0 || board[i-1][j] == '.') && (j==0 || board[i][j-1] == '.') {
				c++
			}
		}
	}
	

	fmt.Println(c)
	return c
}


func main() {
	countBattleships([][]byte{{'X','.','.','X'},{'.','.','.','X'},{'.','.','.','X'}})
}