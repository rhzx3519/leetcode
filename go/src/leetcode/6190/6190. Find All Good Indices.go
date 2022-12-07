/**
@author ZhengHao Lou
Date    2022/09/26
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/find-all-good-indices/
*/
func goodIndices(nums []int, k int) (ans []int) {
	n := len(nums)
	var left, right = make([]int, n), make([]int, n)
	for i := range nums {
		right[i] = i
		left[i] = i
	}
	for i := 0; i < n; {
		j := i + 1
		for j < n && nums[j] <= nums[j-1] {
			left[j] = i
			j++
		}
		i = j
	}

	for i := n - 1; i >= 0; {
		j := i - 1
		for j >= 0 && nums[j] <= nums[j+1] {
			right[j] = i
			j--
		}
		i = j
	}

	//fmt.Println(left)
	//fmt.Println(right)
	for i := k; i < n-k; i++ {
		//fmt.Println(i, left[i], i-k, right[i], i+k)
		if left[i-1] <= i-k && right[i+1] >= i+k {
			ans = append(ans, i)
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	goodIndices([]int{2, 1, 1, 1, 3, 4, 1}, 2)
	goodIndices([]int{2, 1, 1, 2}, 2)
	// 4,5
	goodIndices([]int{878724, 201541, 179099, 98437, 35765, 327555, 475851, 598885, 849470, 943442}, 4)

}
