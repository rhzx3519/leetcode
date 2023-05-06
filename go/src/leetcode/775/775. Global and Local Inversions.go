/*
*
@author ZhengHao Lou
Date    2022/11/16
*/
package main

import "fmt"

func isIdealPermutation(nums []int) bool {
	var a, b int
	for i := range nums {
		if i+1 < len(nums) && nums[i] > nums[i+1] {
			a++
		}
	}

	var merge = func(l, m, r int) (reversed int) {
		var tmp = make([]int, r-l+1)
		var k int
		var i, j = l, m + 1
		for ; i <= m && j <= r; k++ {
			if nums[i] <= nums[j] {
				tmp[k] = nums[i]
				i++
			} else {
				tmp[k] = nums[j]
				reversed += m - i + 1
				j++
			}
		}
		for ; i <= m; i++ {
			tmp[k] = nums[i]
			k++
		}
		for ; j <= r; j++ {
			tmp[k] = nums[j]
			k++
		}
		copy(nums[l:], tmp)
		return
	}

	var mergeSort func(l, r int) int
	mergeSort = func(l, r int) int {
		var tot int
		if l < r {
			m := l + (r-l)>>1
			tot += mergeSort(l, m)
			tot += mergeSort(m+1, r)
			tot += merge(l, m, r)
		}
		return tot
	}

	b = mergeSort(0, len(nums)-1)
	fmt.Println(a, b)
	return a == b
}

func main() {
	isIdealPermutation([]int{1, 0, 2})
	isIdealPermutation([]int{1, 2, 0})
}
