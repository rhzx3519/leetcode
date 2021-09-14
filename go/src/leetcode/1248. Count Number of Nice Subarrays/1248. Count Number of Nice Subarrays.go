package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]&1
	}
	//fmt.Println(sum)
	var mapper = make(map[int]int)
	mapper[0] = 1
	var count int
	for i := 1; i <= n; i++ {
		if t, ok := mapper[sum[i] - k]; ok {
			count += t
		}
		mapper[sum[i]]++
	}

	return count
}

func main() {
	fmt.Println(numberOfSubarrays([]int{1,1,2,1,1}, 3))
	fmt.Println(numberOfSubarrays([]int{2,4,6}, 1))
	fmt.Println(numberOfSubarrays([]int{2,2,2,1,2,2,1,2,2,2}, 2))
}
