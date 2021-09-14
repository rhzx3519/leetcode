package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	mapper := make(map[int]int)
	mapper[0] = -1

	var s int
	for i, num := range nums {
		s += num
		s %= k
		j, ok := mapper[s]
		if !ok {
			mapper[s] = i
		} else {
			if i-j > 1 {
				return true
			}
		}

	}

	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{23,2,4,6,7}, 6))
	fmt.Println(checkSubarraySum([]int{23,2,6,4,7}, 6))
	fmt.Println(checkSubarraySum([]int{23,2,6,4,7}, 13))
	fmt.Println(checkSubarraySum([]int{0,0}, 1))
	fmt.Println(checkSubarraySum([]int{0}, 1))
	fmt.Println(checkSubarraySum([]int{5,0,0}, 3))
}