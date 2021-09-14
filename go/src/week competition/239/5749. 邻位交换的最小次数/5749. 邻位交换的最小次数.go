package main

import (
	"fmt"
)

func getMinSwaps(num string, k int) int {
	n := len(num)
	numsK := []int{}
	nums := []int{}
	for _, c := range num {
		numsK = append(numsK, int(c - '0'))
		nums = append(nums, int(c - '0'))
	}

	for i := 0; i < k; i++ {
		numsK = nextPermutation(numsK)
	}

	var count int
	for i := 0; i < n; i++ {
		if nums[i] == numsK[i] {
			continue
		}

		for j := i+1; j < n; j++ {
			if nums[j] != numsK[i] {
				continue
			}
			//for k := j-1; k >= i; k-- {
			//	count++
			//	nums[k], nums[k+1] = nums[k+1], nums[k]
			//}
			count += j-i
			copy(nums[i+1:j+1], nums[i:j])
			nums[i] = numsK[i]
			break
		}
	}

	return count
}

func nextPermutation(nums []int) []int {
	n := len(nums)
	i := n-2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n-1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}

func main() {
	fmt.Println(getMinSwaps("5489355142", 4))
	fmt.Println(getMinSwaps("11112", 4))
}
