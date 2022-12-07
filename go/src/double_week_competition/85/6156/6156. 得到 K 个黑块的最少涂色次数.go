/**
@author ZhengHao Lou
Date    2022/08/20
*/
package main

import "fmt"

func minimumRecolors(blocks string, k int) int {
	var mx, c int
	for i := range blocks {
		switch blocks[i] {
		case 'B':
			c++
		}

		if i >= k {
			if blocks[i-k] == 'B' {
				c--
			}
		}

		if c > mx {
			mx = c
		}
	}

	if mx >= k {
		return 0
	}
	return k - mx
}

func main() {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("WBWBBBW", 2))
}
