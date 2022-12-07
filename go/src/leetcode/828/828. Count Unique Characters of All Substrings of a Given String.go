/**
@author ZhengHao Lou
Date    2022/09/06
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/
*/
const N = 26

func uniqueLetterString(s string) int {
	st := make([][]int, N)
	for i := range s {
		j := encode(s[i])
		st[j] = append(st[j], i)
	}

	var ans int
	for i := 0; i < N; i++ {
		st[i] = append([]int{-1}, append(st[i], len(s))...)
		for j := 1; j < len(st[i])-1; j++ {
			a, b, c := st[i][j-1], st[i][j], st[i][j+1]
			ans += (b - a) * (c - b)
		}
	}
	fmt.Println(ans)
	return ans
}

func encode(b byte) int {
	return int(b - 'A')
}

func main() {
	uniqueLetterString("ABC")
}
