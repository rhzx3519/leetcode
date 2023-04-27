package main

import (
	"fmt"
)

/*
*
https://leetcode.cn/problems/sliding-subarray-beauty/
*/
const N = 50

func getSubarrayBeauty(nums []int, k int, x int) []int {
	state := make([]int, N+1)
	var ans []int
	for i, t := range nums {
		if t < 0 {
			state[-t]++
		}
		if i >= k && nums[i-k] < 0 {
			state[-nums[i-k]]--
		}
		if i >= k-1 {
			var mi int
			for j, c := N, 0; j >= 1; j-- {
				if state[j] > 0 {
					c += state[j]
				}
				if c >= x {
					mi = -j
					break
				}
			}
			ans = append(ans, mi)
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	//getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2)
	//getSubarrayBeauty([]int{-1, -2, -3, -4, -5}, 2, 2)
	//getSubarrayBeauty([]int{-3, 1, 2, -3, 0, -3}, 2, 1)
	getSubarrayBeauty([]int{-46, -34, -46}, 3, 3)
}
