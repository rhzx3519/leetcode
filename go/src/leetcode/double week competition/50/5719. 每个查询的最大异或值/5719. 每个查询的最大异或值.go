package main

import "fmt"

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	ans := make([]int, n)

	t := 0
	for _, num := range nums {
		t ^= num
	}

	for i := n-1; i >= 0; i-- {
		ans[n-1-i] = int(max(uint32(t), maximumBit))
		t ^= nums[i]
	}


	return ans
}

func max(x uint32, maximumBit int) uint32 {
	return ^x & (1 << maximumBit - 1)
}

func main() {
	//fmt.Printf("%b\n", max(0b100, 2))
	fmt.Println(getMaximumXor([]int{0,1,1,3}, 2))
	fmt.Println(getMaximumXor([]int{2,3,4,7}, 3))
	fmt.Println(getMaximumXor([]int{0,1,2,2,5,7}, 3))
}
