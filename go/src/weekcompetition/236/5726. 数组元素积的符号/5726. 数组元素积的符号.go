package main

import "fmt"

func arraySign(nums []int) int {
	neg := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			neg++
		}
	}

	if neg&1 == 1 {
		return -1
	}
	return 1
}

func main() {
	fmt.Println(arraySign([]int{-1,-2,-3,-4,3,2,1}))
	fmt.Println(arraySign([]int{1,5,0,2,-3}))
	fmt.Println(arraySign([]int{-1,1,-1,1,-1}))
}