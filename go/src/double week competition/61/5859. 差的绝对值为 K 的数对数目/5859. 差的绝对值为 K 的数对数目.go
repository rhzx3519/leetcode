package main

import "fmt"

func countKDifference(nums []int, k int) int {
	var count int
	mapper := make(map[int]int)
	for i := range nums {
		count += mapper[nums[i] + k]
		count += mapper[nums[i] - k]
		mapper[nums[i]]++
	}
	return count
}

func main() {
	fmt.Println(countKDifference([]int{1,2,2,1}, 1))
	fmt.Println(countKDifference([]int{1,3}, 3))
	fmt.Println(countKDifference([]int{3,2,1,5,4}, 2))
}
