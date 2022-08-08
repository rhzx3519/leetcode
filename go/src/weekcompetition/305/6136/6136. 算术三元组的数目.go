/**
@author ZhengHao Lou
Date    2022/08/07
*/
package main

func arithmeticTriplets(nums []int, diff int) int {
	n := len(nums)
	var left, right = make([]int, n), make([]int, n)
	counter := make(map[int]int)
	for i := range nums {
		left[i] = counter[nums[i]-diff]
		counter[nums[i]]++
	}
	counter = make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		right[i] = counter[nums[i]+diff]
		counter[nums[i]]++
	}
	var ans int
	for i := 0; i < n; i++ {
		ans += left[i] * right[i]
	}
	return ans
}


