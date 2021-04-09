package main

import "fmt"

func removeDuplicates(nums []int) int {
	last := 0
	count := 1
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[last] {
			if count >= 2 {
				continue
			} else {
				count++
				last++
				nums[last] = nums[i]
			}
		} else {
			last++
			nums[last] = nums[i]
			count = 1
		}
	}
	fmt.Println(nums[:last+1])
	return last
}

func main() {
	removeDuplicates([]int{1,1,1,2,2,3})
	removeDuplicates([]int{0,0,1,1,1,1,2,3,3})
}