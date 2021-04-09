package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mem := make(map[int]int)
	for i, num := range nums {
		other := target - num
		if _, ok := mem[other]; ok {
			return []int{mem[other], i}
		}
		mem[num] = i
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
