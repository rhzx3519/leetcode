package main

import (
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/minimum-seconds-to-equalize-a-circular-array/?envType=daily-question&envId=2024-01-30
思路：扩散问题，nums的元素必然变为其中的一个元素，每个元素之间的最大距离决定了所需的秒数，所有相同元素的最大距离的最小值为所需最小秒数
*/
func minimumSeconds(nums []int) int {
	mem := make(map[int]int)
	mx := make(map[int]int)

	nums = append(nums, nums...)
	for i, x := range nums {
		if j, ok := mem[x]; ok {
			mx[x] = max(mx[x], i-j-1)
		}
		mem[x] = i
	}
	var mi = math.MaxInt32
	for _, x := range mx {
		mi = min(mi, x)
	}
	return (mi + 1) >> 1
}

func main() {
	fmt.Println(minimumSeconds([]int{1, 2, 1, 2}))
	fmt.Println(minimumSeconds([]int{2, 1, 3, 3, 2}))
	fmt.Println(minimumSeconds([]int{5, 5, 5, 5}))
}
