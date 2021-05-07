package main

import "fmt"

func getXORSum(arr1 []int, arr2 []int) int {
	var x1, x2 int
	for _, t := range arr2 {
		x2 ^= t
	}
	for _, t := range arr1 {
		x1 ^= t &x2
	}
	return x1
}

func main() {
	//fmt.Printf("%b\n", getXORSum([]int{1,2,3}, []int{6,5}))
	fmt.Printf("%b\n", getXORSum([]int{12}, []int{4}))
}
