package main

import "fmt"

/*
*
https://leetcode.cn/problems/count-sorted-vowel-strings/
f[i][j] 表示以j结尾，且长度等于i的按字典序排列的字符串数量
f[i+1][j] = sum(f[i][0-j])
*/
const N = 5

func countVowelStrings(n int) (tot int) {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, N)
	}
	f[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k <= j; k++ {
				f[i][j] += f[i-1][k]
			}
		}
	}
	for i := 0; i < N; i++ {
		tot += f[n][i]
	}
	fmt.Println(tot)
	return
}

func main() {
	countVowelStrings(33)
}
