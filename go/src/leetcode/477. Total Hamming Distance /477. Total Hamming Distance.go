package main

import "fmt"

const N = 30

func totalHammingDistance(nums []int) int {
	n := len(nums)
	var ans int
	for i := 0; i <= N; i++ {
		var count1 int
		for j := range nums {
			count1 += nums[j]>>i&1
		}
		ans += count1*(n-count1)
	}
	return ans
}

func hamming(a, b int) int {
	var count int
	for i := N; i >= 0; i-- {
		count += (a>>i&1) ^ (b>>i&1)
	}
	return count
}

func main() {
	fmt.Println(totalHammingDistance([]int{4,14,2}))
	fmt.Println(totalHammingDistance([]int{4,14,4}))
}
