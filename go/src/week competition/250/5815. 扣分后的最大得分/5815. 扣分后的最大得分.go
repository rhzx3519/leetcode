package main

import "fmt"

/**
思路: dp
row[i](0<=i<=n-1), 表示选择i列时能获得的最大分数
 */
func maxPoints(points [][]int) int64 {
	var score = int64(0)
	m, n := len(points), len(points[0])

	left := make([]int64, n)	// left[i]表示[0, i]中能获取的最大得分
	right := make([]int64, n)	// right[i]表示[i, n-1]中能获取的最大得分
	row := make([]int64, n)

	for r := 0; r < m; r++ {
		left[0] = row[0]
		for i := 1; i < n; i++ {
			left[i] = max(row[i], left[i-1] - 1)
		}
		right[n-1] = row[n-1]
		for i := n-2; i >= 0; i-- {
			right[i] = max(row[i], right[i+1] - 1)
		}

		for i := 0; i < n; i++ {
			row[i] = int64(points[r][i]) + max(left[i], right[i])
		}
	}

	for i := 0; i < n; i++ {
		score = max(score, row[i])
	}
	fmt.Println(score)
	return score
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxPoints([][]int{{1,2,3},{1,5,1},{3,1,1}})
	maxPoints([][]int{{1,5},{2,3},{4,2}})
}