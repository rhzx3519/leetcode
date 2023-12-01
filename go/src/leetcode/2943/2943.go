package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/maximize-area-of-square-hole-in-grid/
*/
func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)

	h, v := consecutive(hBars), consecutive(vBars)
	fmt.Println(h, v)
	t := min(h, v) + 1
	return t * t
}

func consecutive(arr []int) int {
	var mx, h int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+1 != arr[i+1] {
			h = 0
		} else {
			h++
		}
		mx = max(mx, h)
	}
	return mx + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//maximizeSquareHoleArea(2, 1, []int{2, 3}, []int{2})
	//maximizeSquareHoleArea(1, 1, []int{2}, []int{2})
	//maximizeSquareHoleArea(2, 3, []int{2, 3}, []int{2, 4})
	maximizeSquareHoleArea(4, 40, []int{5, 3, 2, 4}, []int{6, 33, 34, 36, 41})
}
