package main

import "fmt"

func isMonotonic(A []int) bool {
	n := len(A)
	if n <= 2 {
		return true
	}
	dir := 0
	for i := 0; i < n-1; i++ {
		if A[i] < A[i+1] && dir==-1 {
			return false
		} else if A[i] > A[i+1] && dir==1 {
			return false
		}
		if dir == 0 {
			if A[i] < A[i+1] {
				dir = 1
			} else if A[i] > A[i+1] {
				dir = -1
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isMonotonic([]int{6,5,4,4}))
	fmt.Println(isMonotonic([]int{1,2,2,3}))
	fmt.Println(isMonotonic([]int{1,3,2}))
	fmt.Println(isMonotonic([]int{1,2,4,5}))
	fmt.Println(isMonotonic([]int{1,1,1}))
}
