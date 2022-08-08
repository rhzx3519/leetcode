package main

import (
	"fmt"
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		if arr[i] - arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[n-1]
}

func main() {
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{2,2,1,2,1}))
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{100,1,1000}))
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{1,2,3,4,5}))
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{73,98,9}))
}