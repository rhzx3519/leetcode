/**
@author ZhengHao Lou
Date    2022/09/12
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/optimal-partition-of-string/
*/
func partitionString(s string) int {
	n := len(s)
	var counter = make([]int, 26)
	for i := range counter {
		counter[i] = -1
	}
	next := make([]int, n)
	for i := range s {
		next[i] = n
		if j := counter[int(s[i]-'a')]; j != -1 {
			next[j] = i
		}
		counter[int(s[i]-'a')] = i
	}

	fmt.Println(next)
	var c = 1
	var r = next[0]
	for i := range s {
		r = min(r, next[i])
		if r == i {
			r = next[i]
			c++
		}
	}

	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(partitionString("abacaba"))
	fmt.Println(partitionString("ssssss"))
	fmt.Println(partitionString("shkqbyutdvknyrpjof")) // 2
}
