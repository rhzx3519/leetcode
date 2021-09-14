package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	for i := range citations {
		if citations[i] >= n-i {
			return n-i
		}
	}
	return 0
}

func main() {
	fmt.Println(hIndex([]int{3,0,6,1,5}))
}
