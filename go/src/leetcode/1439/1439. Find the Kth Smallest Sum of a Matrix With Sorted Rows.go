package main

import (
	"fmt"
	"sort"
)

func kthSmallest(mat [][]int, k int) int {
	arr := []int{0}
	for _, row := range mat {
		var tmp []int
		for j := range row {
			for i := range arr {
				tmp = append(tmp, arr[i]+row[j])
			}
		}
		sort.Ints(tmp)
		arr = tmp[:min(k, len(tmp))]
	}
	return arr[k-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(kthSmallest([][]int{{1, 3, 11}, {2, 4, 6}}, 5))
	//fmt.Println(kthSmallest([][]int{{1, 3, 11}, {2, 4, 6}}, 9))
	//fmt.Println(kthSmallest([][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}, 7))
	fmt.Println(kthSmallest([][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}, 14))
}
