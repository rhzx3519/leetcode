package main

import "fmt"

func chalkReplacer(chalk []int, k int) int {
	var s int
	for _, c := range chalk {
		s += c
	}
	k %= s
	for i, c := range chalk {
		k -= c
		if k < 0 {
			return i
		}
	}

	return len(chalk) - 1
}

func main() {
	fmt.Println(chalkReplacer([]int{5,1,5}, 22))
	fmt.Println(chalkReplacer([]int{3,4,1,2}, 25))
}
