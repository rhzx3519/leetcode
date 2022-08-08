/**
@author ZhengHao Lou
Date    2021/10/27
*/
package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	degree := make([][]int, n + 1)
	for i := range degree {
		degree[i] = []int{0, 0}
	}


	for i := range trust {
		a, b := trust[i][0], trust[i][1]
		degree[a][1]++
		degree[b][0]++
	}
	for i := 1; i <= n; i++ {
		if degree[i][0] == n - 1 && degree[i][1] == 0 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1,2}}))
	fmt.Println(findJudge(3, [][]int{{1,3},{2,3}}))
	fmt.Println(findJudge(3, [][]int{{1,3},{2,3},{3,1}}))
	fmt.Println(findJudge(3, [][]int{{1,2},{2,3}}))
	fmt.Println(findJudge(4, [][]int{{1,3},{1,4},{2,3},{2,4},{4,3}}))
}
