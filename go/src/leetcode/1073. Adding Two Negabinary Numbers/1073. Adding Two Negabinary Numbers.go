/**
@author ZhengHao Lou
Date    2021/11/12
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/adding-two-negabinary-numbers/
思路：arr[i]进位，arr[i+1], arr[i+1] 也要进位
 */
func addNegabinary(arr1 []int, arr2 []int) []int {
	n1, n2 := len(arr1), len(arr2)
	if n1 > n2 {
		return addNegabinary(arr2, arr1)
	}
	var N = n2 + 4
	arr := make([]int, N)
	for i := n1 - 1; i >= 0; i-- {
		arr[n1 - 1 - i] += arr1[i]
	}
	for i := n2 - 1; i >= 0; i-- {
		arr[n2 - 1 - i] += arr2[i]
	}

	for i := 0; i + 2 < N; i++ {
		var c = arr[i] >> 1
		arr[i] &= 1
		arr[i+1] += c
		arr[i+2] += c
	}
	var k = N - 3
	for ; k > 0 && arr[k] == 0; k-- {
	}
	arr = arr[:k + 1]

	reverse(arr)
	return arr
}

func reverse(arr []int) {
	l, r := 0, len(arr) - 1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func main() {
	fmt.Println(addNegabinary([]int{1}, []int{1}))
	//fmt.Println(addNegabinary([]int{1,1}, []int{1}))
	//fmt.Println(addNegabinary([]int{1,1,1,1,1}, []int{1,0,1}))
}
