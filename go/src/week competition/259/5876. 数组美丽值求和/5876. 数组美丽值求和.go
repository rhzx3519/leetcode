package main

import "fmt"

func sumOfBeauties(nums []int) int {
	n := len(nums)

	c1 := make([]bool, n)
	c2 := make([]bool, n)
	var maxVal = nums[0]
	for i := 1; i < n-1; i++ {
		c2[i] = nums[i] > maxVal
		maxVal = max(maxVal, nums[i])
	}

	var minVal = nums[n-1]
	for i := n-2; i >= 1; i-- {
		c1[i] = nums[i] < minVal
		minVal = min(minVal, nums[i])
	}

	var score int
	for i := 1; i < n-1; i++ {
		if c1[i] && c2[i] {
			score += 2
		} else if nums[i] > nums[i-1] && nums[i] < nums[i+1] {
			score += 1
		}
	}
	return score
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(sumOfBeauties([]int{1,2,3}))
	fmt.Println(sumOfBeauties([]int{2,4,6,4}))
	fmt.Println(sumOfBeauties([]int{3,2,1}))
}