package main

import "fmt"

func numSubarraysWithSum(nums []int, goal int) int {
	mapper := make(map[int]int)
	mapper[0] = 1
	var ans, s int
	for _, d := range nums {
		s += d
		if c, ok := mapper[s - goal]; ok {
			ans += c
		}
		mapper[s]++
	}
	return ans
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0,0,0,0,0}, 0))
}

