package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-special-quadruplets/
 */
func countQuadruplets(nums []int) int {
	var count int

	for d := range nums {
		for a := 0; a < d; a++ {
			for b := a+1; b < d; b++ {
				for c := b+1; c < d; c++ {
					if nums[a] + nums[b] + nums[c] == nums[d] {
						count++
					}
				}
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countQuadruplets([]int{1,2,3,6}))
	fmt.Println(countQuadruplets([]int{3,3,6,4,5}))
	fmt.Println(countQuadruplets([]int{1,1,1,3,5}))
}
