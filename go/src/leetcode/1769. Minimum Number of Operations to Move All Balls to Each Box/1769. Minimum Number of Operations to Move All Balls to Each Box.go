package main

import "fmt"

func minOperations(boxes string) []int {
	n := len(boxes)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if boxes[j] == '1' {
				ans[i] += abs(j - i)
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minOperations("110"))
	fmt.Println(minOperations("001011"))
}
