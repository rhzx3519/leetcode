/**
@author ZhengHao Lou
Date    2022/04/02
*/
package main

import "fmt"

func sumScores(s string) int64 {
	nxt := getNext(s)
	fmt.Println(nxt)

	var c int64
	return c
}

func getNext(p string) []int {
	n := len(p)
	nxt := make([]int, n)
	for x := n - 1; x >= 0; x-- {
		var i = x
		for ; i >= 0; i-- {
			if p[0:i] == p[x-i+1:x+1] {
				break
			}
		}
		nxt[x] = i
	}
	return nxt
}

func main() {
	sumScores("babab")
}
