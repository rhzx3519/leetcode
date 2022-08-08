/**
@author ZhengHao Lou
Date    2021/12/30
*/
package main

import (
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)
	if n%groupSize != 0 {
		return false
	}

	counter := make(map[int]int)
	for _, num := range hand {
		counter[num]++
	}
	sort.Ints(hand)
	for _, num := range hand {
		if counter[num] == 0 {
			continue
		}
		for i := num; i < num + groupSize; i++ {
			if counter[i] == 0 {
				return false
			}
			counter[i]--
		}
	}
	return true
}

func main() {
	fmt.Println(isNStraightHand([]int{1,2,3,6,2,3,4,7,8}, 3))
}