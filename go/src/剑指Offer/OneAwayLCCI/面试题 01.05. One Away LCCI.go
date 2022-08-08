/**
@author ZhengHao Lou
Date    2022/05/13
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/one-away-lcci/
*/
func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		f[i][0] = i
	}
	for i := 1; i <= n; i++ {
		f[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if first[i-1] == second[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j-1], min(f[i-1][j], f[i][j-1])) + 1
			}
		}
	}

	fmt.Println(f)
	return f[m][n] <= 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	oneEditAway("pale", "ple")
	oneEditAway("pales", "pal")
}
