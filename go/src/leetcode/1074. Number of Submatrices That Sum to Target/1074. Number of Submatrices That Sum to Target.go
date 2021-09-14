package main

import "fmt"

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	var count int
	for i := 0; i < m; i++ {	// 下界
		var sum = make([]int, n)
		for j := i; j < m; j++ {  // 上界
			for c := 0; c < n; c++ { // 每列
				sum[c] += matrix[j][c]
			}
			count += sub(sum, target)
		}
	}

	return count
}

func sub(sum []int, k int) int {
	mapper := make(map[int]int)
	mapper[0] = 1

	var count, pre int
	for _, num := range sum {
		pre += num
		if n, ok := mapper[pre - k]; ok {
			count += n
		}
		mapper[pre]++
	}
	return count
}

func main() {
	fmt.Println(numSubmatrixSumTarget([][]int{{0,1,0},{1,1,1},{0,1,0}}, 0))
}
