package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	ans := []int{}
	l1, r1, l2, r2 := 0, m-1, 0, n-1
	for l1<=r1 && l2<=r2 {
		i, j := l1, l2
		for ; j <= r2; j++ {
			ans = append(ans, matrix[i][j])
		}
		i++
		j--
		for ; i <= r1; i++ {
			ans = append(ans, matrix[i][j])
		}
		i--
		j--

		for ; l1 != r1 && j > l2; j-- {
			ans = append(ans, matrix[i][j])
		}

		for ; l2 != r2 && i > l1; i-- {
			ans = append(ans, matrix[i][j])
		}
		l1++
		l2++
		r1--
		r2--
	}
	return ans
}

func main() {
	fmt.Println(spiralOrder([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}))
}
