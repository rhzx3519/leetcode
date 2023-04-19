package main

import (
	"fmt"
	"strings"
)

/*
*
https://leetcode.cn/problems/substring-xor-queries/
*/
func substringXorQueries(s string, queries [][]int) [][]int {
	var ans = make([][]int, len(queries))
	for i := range queries {
		val := queries[i][0] ^ queries[i][1]
		sub := tob(val)
		if j := strings.Index(s, sub); j == -1 {
			ans[i] = []int{-1, -1}
		} else {
			ans[i] = []int{j, j + len(sub) - 1}
		}
	}
	fmt.Println(ans)
	return ans
}

func tob(x int) string {
	if x == 0 {
		return "0"
	}
	var bytes []byte
	for ; x != 0; x >>= 1 {
		bytes = append(bytes, byte('0'+x&1))
	}
	for l, r := 0, len(bytes)-1; l < r; l, r = l+1, r-1 {
		bytes[l], bytes[r] = bytes[r], bytes[l]
	}
	return string(bytes)
}

func main() {
	substringXorQueries("101101", [][]int{{0, 5}, {1, 2}})
	substringXorQueries("0101", [][]int{{12, 8}})
	substringXorQueries("1", [][]int{{4, 5}})
	substringXorQueries("111010110", [][]int{{4, 2}, {3, 3}, {6, 4}, {9, 9}, {10, 28}, {0, 470}, {5, 83}, {10, 28}, {8, 15}, {6, 464}, {0, 3}, {5, 8}, {7, 7}, {8, 8}, {6, 208}, {9, 15}, {2, 2}, {9, 95}})
}
