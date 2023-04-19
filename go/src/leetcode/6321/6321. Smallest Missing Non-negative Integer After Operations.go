package main

import "fmt"

/*
*
https://leetcode.cn/problems/smallest-missing-non-negative-integer-after-operations/
(b - a) % d == 0
*/
func findSmallestInteger(nums []int, value int) int {
	cnt := map[int]int{}
	for i := range nums {
		nums[i] = minPositve(nums[i], value)
		cnt[nums[i]]++
	}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		x := minPositve(i-value, value)
		if cnt[x] == 0 {
			return i
		}
		cnt[x]--
	}
	return len(nums)
}

func minPositve(a, b int) int {
	return (a + abs(a/b*b) + b) % b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 7))
	fmt.Println(findSmallestInteger([]int{3, 0, 3, 2, 4, 2, 1, 1, 0, 4}, 5))          // 10
	fmt.Println(findSmallestInteger([]int{3, 2, 3, 1, 0, 1, 4, 2, 3, 1, 4, 1, 3}, 5)) // 5
}
