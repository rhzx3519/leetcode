package main

import (
	"fmt"
	"sort"
)

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	stones := make([]int, n)
	for i, _ := range stones {
		stones[i] = i
	}

	sort.Slice(stones, func(i, j int) bool {
		return aliceValues[stones[i]] + bobValues[stones[i]] > aliceValues[stones[j]] + bobValues[stones[j]]
	})

	a, b := 0, 0
	for i := 0; i < n; i++ {
		if (i&1) == 0 {
			a += aliceValues[stones[i]]
		} else {
			b += bobValues[stones[i]]
		}
	}

	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}

}

func main() {
	aliceValues := []int{9,9,5,5,2,8,2,4,10,2,3,3,4}
	bobValues := []int{9,5,3,4,4,6,6,6,4,3,7,5,10}
	fmt.Println(stoneGameVI(aliceValues,  bobValues))


	a := []int{0,1,2}
	b := map[int]int{0:1, 1:3, 2:2}
	sort.Slice(a, func(i, j int) bool {
		return b[a[i]] > b[a[j]]
	})
	// expected: {1,2,0}
	// actual: {1,0,2}
	fmt.Println(a)
}
