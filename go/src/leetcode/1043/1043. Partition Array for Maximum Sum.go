package main

import "fmt"

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	f := make([]int, n+1)
	for i := range arr {
		for j, mx := i, 0; j >= 0 && j > i-k; j-- {
			mx = max(mx, arr[j])
			f[i+1] = max(f[i+1], f[j]+(i-j+1)*mx)
		}
	}
	fmt.Println(f)
	return f[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
}
