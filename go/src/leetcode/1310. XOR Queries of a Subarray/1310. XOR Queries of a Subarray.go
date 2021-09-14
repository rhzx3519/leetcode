package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	pre := make([]int, len(arr)+1)
	for i := 1; i <= len(arr); i++ {
		pre[i] = pre[i-1]^arr[i-1]
	}
	for i, q := range queries {
		l, r := q[0], q[1]
		ans[i] = pre[r+1] ^ pre[l]
	}
	return ans
}

func main() {
	fmt.Println(xorQueries([]int{1,3,4,8}, [][]int{{0,1},{1,2},{0,3},{3,3}}))
	fmt.Println(xorQueries([]int{4,8,2,10}, [][]int{{2,3},{1,3},{0,0},{0,3}}))
}