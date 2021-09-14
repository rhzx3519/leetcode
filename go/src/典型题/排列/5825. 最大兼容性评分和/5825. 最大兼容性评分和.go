package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-compatibility-score-sum/
思路：回溯，排列
 */
func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m, _ := len(students), len(students[0])

	scores := make([][]int, m)
	for i := 0; i < m; i++ {
		scores[i] = make([]int, m)
		for j := 0; j < m; j++ {
			scores[i][j] = getScore(students[i], mentors[j])
		}
	}
	visited := make([]bool, m)

	var ans int
	var backtrace func(i, k, score int)
	backtrace = func(i, k, score int) {
		if k == 0 {
			ans = max(ans, score)
			return
		}

		for j := 0; j < m; j++ {
			if visited[j] {
				continue
			}
			visited[j] = true
			backtrace(i+1, k-1, score + scores[i][j])
			visited[j] = false
		}

	}

	backtrace(0, m, 0)
	return ans
}

func getScore(student, mentor []int) int {
	var score int
	for i := range student {
		if student[i] == mentor[i] {
			score++
		}
	}
	return score
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 8
	fmt.Println(maxCompatibilitySum([][]int{{1,1,0},{1,0,1},{0,0,1}}, [][]int{{1,0,0},{0,0,1},{1,1,0}}))
	// 0
	fmt.Println(maxCompatibilitySum([][]int{{0,0},{0,0},{0,0}}, [][]int{{1,1},{1,1},{1,1}}))
	// 9
	fmt.Println(maxCompatibilitySum([][]int{{1,1,1},{0,0,1},{0,0,1},{0,1,0}}, [][]int{{1,0,1},{0,1,1},{0,1,0},{1,1,0}}))
	// 24
	fmt.Println(maxCompatibilitySum([][]int{{0,0,1,1,1,0,1},{0,1,1,0,0,0,0},{0,0,1,1,1,1,1},{0,1,0,0,1,0,1},{1,0,1,1,1,1,1}}, [][]int{{0,1,1,0,0,0,0},{0,1,0,0,0,0,1},{0,1,0,1,0,0,1},{1,0,0,0,1,0,1},{1,1,1,1,1,0,0}}))
}

//[[0,0,1,1,1,0,1],[0,1,1,0,0,0,0],[0,0,1,1,1,1,1],[0,1,0,0,1,0,1],[1,0,1,1,1,1,1]]
//[[0,1,1,0,0,0,0],[0,1,0,0,0,0,1],[0,1,0,1,0,0,1],[1,0,0,0,1,0,1],[1,1,1,1,1,0,0]]
//24