/**
思路：nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
time: O(n), space: O(n)
 */
package main

import (
	"fmt"
	"math"
	"strconv"
)

var MOD = int(math.Pow10(9)) + 7

func countNicePairs(nums []int) int {
	count := 0
	//n := len(nums)
	mapper := make(map[int]int)
	for _, num := range nums {
		mapper[num - rev(num)]++
	}

	for _, n := range mapper {
		count = (count + sum(n)) % MOD
	}

	return count
}

func sum(n int) int {
	s := 0
	for i := 1; i < n; i++ {
		s += i
	}
	return s
}

func rev(num int) int {
	s := strconv.Itoa(num)
	runes := []rune(s)
	l, r := 0, len(s)-1
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	s = string(runes)
	i, _ :=strconv.Atoi(s)
	return i
}

func main() {
	fmt.Println(countNicePairs([]int{42,11,1,97}))
	fmt.Println(countNicePairs([]int{13,10,35,24,76}))
}