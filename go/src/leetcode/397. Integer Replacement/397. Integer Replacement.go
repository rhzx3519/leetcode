/**
@author ZhengHao Lou
Date    2021/11/19
*/
package main

import "fmt"

func integerReplacement(n int) int {
	var step int
	vis := make(map[int]int)
	vis[n]++

	que := []int{n}
	for len(que) != 0 {
		sz := len(que)
		for i := 0; i < sz; i++ {
			n = que[0]
			que = que[1:]
			if n == 1 {
				return step
			}
			if n&1 == 0 {
				if vis[n>>1] == 0 {
					vis[n>>1]++
					que = append(que, n>>1)
				}
			} else {
				if vis[n+1] == 0 {
					vis[n+1]++
					que = append(que, n+1)
				}
				if vis[n-1] == 0 {
					vis[n-1]++
					que = append(que, n-1)
				}
			}
		}

		step++
	}
	return 0
}

func main() {
	fmt.Println(integerReplacement(8))
	fmt.Println(integerReplacement(7))
	fmt.Println(integerReplacement(4))
}
