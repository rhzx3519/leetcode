/**
@author ZhengHao Lou
Date    2022/07/05
*/
package main

/**
https://leetcode.cn/problems/spiral-matrix-iv/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	var r1, r2 = 0, m - 1
	var c1, c2 = 0, n - 1
	var val int
	for r1 <= r2 && c1 <= c2 {
		var i, j = r1, c1
		for ; j <= c2; j++ {
			head, val = next(head)
			matrix[r1][j] = val
		}
		j--
		i++
		for ; i <= r2; i++ {
			head, val = next(head)
			matrix[i][c2] = val
		}
		i--
		j--
		if r1 != r2 {
			for ; j >= c1; j-- {
				head, val = next(head)
				matrix[r2][j] = val
			}
			j++
		}
		i--
		if c1 != c2 {
			for ; i > r1; i-- {
				head, val = next(head)
				matrix[i][c1] = val
			}
		}

		r1++
		r2--
		c1++
		c2--
	}
	return matrix
}

func next(head *ListNode) (*ListNode, int) {
	if head != nil {
		return head.Next, head.Val
	}
	return nil, -1
}

func main() {
}
