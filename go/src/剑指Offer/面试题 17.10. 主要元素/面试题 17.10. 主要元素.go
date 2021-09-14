package main

import "fmt"

func majorityElement(nums []int) int {
	var x, c = -1, 0
	for _, num := range nums {
		if c == 0 {
			x = num
			c = 1
		} else {
			c += sign(x==num)
		}
	}
	c = 0
	for _, num := range nums {
		if num == x {
			c++
		}
	}
	if c > len(nums)>>1 {
		return x
	}
	return -1
}

func sign(b bool) int {
	if b {
		return 1
	}
	return -1
}

func main() {
	fmt.Println(majorityElement([]int{1,2,5,9,5,9,5,5,5}))
	fmt.Println(majorityElement([]int{3,2}))
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
	fmt.Println(majorityElement([]int{6,5,5}))
	fmt.Println(majorityElement([]int{3,2,3}))
}
