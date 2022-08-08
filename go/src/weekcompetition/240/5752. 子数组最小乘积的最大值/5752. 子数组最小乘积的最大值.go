package main

import (
	"fmt"
	"math"
)

var MOD = int64(math.Pow10(9) + 7)

func maxSumMinProduct(nums []int) int {
	maxVal := int64(0)

	n := len(nums)
	pre := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + nums[i-1]
	}
	var minVal int
	var s int
	for i := 0; i < n; i++ {
		minVal = nums[i]
		s = nums[i]
		for j := i+1; j < n; j++ {
			if minVal > nums[j]	 {
				minVal = nums[j]
			}
			//s = pre[j+1] - pre[i]
			s += nums[j]
			if int64(s)*int64(minVal) > maxVal {
				maxVal = int64(s)*int64(minVal)
			}
		}
	}

	return int(maxVal%MOD)
}

func main() {
	//fmt.Println(maxSumMinProduct([]int{1,2,3,2}))
	//fmt.Println(maxSumMinProduct([]int{2,3,3,1,2}))
	//fmt.Println(maxSumMinProduct([]int{3,1,5,6,4,2}))
	fmt.Println(maxSumMinProduct([]int{1,1,3,2,2,2,1,5,1,5}))
}
