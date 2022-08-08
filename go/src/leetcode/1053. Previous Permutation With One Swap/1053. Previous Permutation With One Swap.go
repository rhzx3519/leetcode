/**
@author ZhengHao Lou
Date    2021/11/10
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/previous-permutation-with-one-swap/
思路：https://leetcode-cn.com/problems/previous-permutation-with-one-swap/solution/jie-ti-si-lu-by-guojian/
 */
func prevPermOpt1(arr []int) []int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {	// 第一次升序的位置
			for j := n - 1; j > i-1; j-- {	// 寻找最右侧第一个小寻找最右侧第一个小于arr于arr[i-1]的元素，且arr[j] != arr[j-1]
				if arr[j] < arr[i-1] && arr[j] != arr[j-1] {
					arr[i-1], arr[j] = arr[j], arr[i-1]
					return arr
				}
			}
		}
	}

	return arr
}

func main() {
	fmt.Println(prevPermOpt1([]int{3,2,1}))
	fmt.Println(prevPermOpt1([]int{1,1,5}))
	fmt.Println(prevPermOpt1([]int{1,9,4,6,7}))
	fmt.Println(prevPermOpt1([]int{3,1,1,3}))
}
