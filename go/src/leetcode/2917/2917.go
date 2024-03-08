package main

import "fmt"

func findKOr(nums []int, k int) int {
	var kor int
	for i := 0; i < 30; i++ {
		var c int
		for _, x := range nums {
			if x>>i&1 == 1 {
				c++
			}
		}
		if c >= k {
			kor |= 1 << i
		}
	}
	return kor
}

func main() {
	fmt.Println(findKOr([]int{7, 12, 9, 8, 9, 15}, 4))
}
