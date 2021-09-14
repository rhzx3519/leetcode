package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + arr[i-1]
	}

	var s int
	for l := 1; l <= n; l += 2 {
		for i := 0; i + l <= n; i++ {
			j := i + l
			s += sum[j] - sum[i]
		}
	}
	return s
}

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1,4,2,5,3}))
	fmt.Println(sumOddLengthSubarrays([]int{1,2}))
	fmt.Println(sumOddLengthSubarrays([]int{10,11,12}))
}
