package main

import "fmt"

/**
Q: 给定一个由数字组成的数组，求当前排列开始的第K个排列
 */
func nextKPermutation(nums []int, k int) []int {
	for i := 0; i < k; i++ {
		nextPermutation(nums)
	}
	return nums
}

func nextPermutation(nums []int)  {
	n := len(nums)
	i := n-2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n-1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
func main() {
	fmt.Println(nextKPermutation([]int{5,4,8,9,3,5,5,1,4,2}, 1))
	//fmt.Println(nextKPermutation([]int{}, ))
	//fmt.Println(nextKPermutation([]int{}, ))
}