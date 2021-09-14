package main

import (
	"fmt"
	"sort"
)

func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] != ranges[j][0] {
			return ranges[i][0] < ranges[j][0]
		}
		return ranges[i][1] < ranges[j][1]
	})
	for _, r := range ranges {
		if right < r[0] || left > r[1] {
			continue
		}
		if left < r[0] {
			return false
		}
		if right < r[1] {
			return true
		}
		left = r[1] + 1
	}
	return right - left + 1 <= 0
}

func main() {
	fmt.Println(isCovered([][]int{{1,2},{3,4},{5,6}}, 2, 5))
	fmt.Println(isCovered([][]int{{1,10},{10,20}}, 21, 21))
	fmt.Println(isCovered([][]int{{1,50}}, 1, 50))
	fmt.Println(isCovered([][]int{{50,50}}, 1, 49))
}

//[[1,50]]
//1
//50

//[[50,50]]
//1
//49