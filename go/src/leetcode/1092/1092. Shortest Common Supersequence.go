package main

import "fmt"

/*
*
https://leetcode.cn/problems/shortest-common-supersequence/
*/
func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	// 求最长公共子序列
	for i := range str1 {
		for j := range str2 {
			if str1[i] == str2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}

	var bytes []byte
	for i, j := m-1, n-1; i >= 0 || j >= 0; {
		var ch byte
		if i >= 0 && j >= 0 {
			if str1[i] == str2[j] {
				ch = str1[i]
				i--
				j--
			} else if f[i+1][j+1] == f[i][j+1] {
				ch = str1[i]
				i--
			} else {
				ch = str2[j]
				j--
			}
		} else if i >= 0 {
			ch = str1[i]
			i--
		} else {
			ch = str2[j]
			j--
		}
		bytes = append(bytes, ch)
	}

	for l, r := 0, len(bytes)-1; l < r; l, r = l+1, r-1 {
		bytes[l], bytes[r] = bytes[r], bytes[l]
	}

	fmt.Println(string(bytes))
	return string(bytes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	shortestCommonSupersequence("abac", "cab")

}
