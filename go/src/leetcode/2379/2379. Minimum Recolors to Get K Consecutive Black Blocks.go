package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/description/
*/
func minimumRecolors(blocks string, k int) int {
	var ans = len(blocks)
	var w int
	for i := range blocks {
		if blocks[i] == 'W' {
			w++
		}
		if i >= k {
			if blocks[i-k] == 'W' {
				w--
			}
		}
		if i >= k-1 {
			ans = min(ans, w)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("WBWW", 2))
}
