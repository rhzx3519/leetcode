package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	var ships int
	sort.Ints(people)
	l, r := 0, len(people) - 1
	for l <= r {
		w := people[l] + people[r]
		if w <= limit {
			l++
		}
		r--
		ships++
	}
	fmt.Println(ships)
	return ships
}

func main() {
	numRescueBoats([]int{1}, 3)
	numRescueBoats([]int{1,2}, 3)
	numRescueBoats([]int{3,2,2,1}, 3)
	numRescueBoats([]int{3,5,3,4}, 5)
	numRescueBoats([]int{5,1,4,2}, 6)
}
