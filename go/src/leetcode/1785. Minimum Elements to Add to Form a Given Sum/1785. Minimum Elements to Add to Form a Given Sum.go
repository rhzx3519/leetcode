package main

import "fmt"

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if goal == sum {
		return 0
	}
	count := 0
	x := goal - sum
	if x < 0 {
		x = -x
	}
	count += x / limit
	if x % limit != 0 {
		count++
	}

	return count
}

func main() {
	fmt.Println(minElements([]int{1,-1,1}, 3, -4))
	fmt.Println(minElements([]int{1,-10,9,1}, 100, 0))
	fmt.Println(minElements([]int{-1,0,1,1,1}, 1, 771843707))
}