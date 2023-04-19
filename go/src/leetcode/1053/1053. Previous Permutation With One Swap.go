/*
*
@author ZhengHao Lou
Date    2021/11/10
*/
package main

import "fmt"

/*
*
https://leetcode-cn.com/problems/previous-permutation-with-one-swap/
*/
func prevPermOpt1(arr []int) []int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		if arr[i-1] > arr[i] { // arr[j] < arr[i-1] > arr[i]
			for j := n - 1; j >= i; j-- {
				if arr[j] < arr[i-1] && arr[j] != arr[j-1] {
					arr[j], arr[i-1] = arr[i-1], arr[j]
					return arr
				}
			}
		}
	}

	return arr
}

func main() {
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7}))
	fmt.Println(prevPermOpt1([]int{3, 1, 1, 3}))
}
