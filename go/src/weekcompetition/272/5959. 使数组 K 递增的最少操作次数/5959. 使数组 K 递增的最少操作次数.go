/**
@author ZhengHao Lou
Date    2021/12/19
*/
package main

import "fmt"

/**
思路：最少操作次数等于数组长度减去最长递增子序列长度
 */
func kIncreasing(arr []int, k int) int {
	var ans int
	n := len(arr)
	for i := 0; i < k; i++ {
		nums := []int{}
		for j := i; j < n; j += k {
			nums = append(nums, arr[j])
		}
		ans += len(nums) - longestIncreasingSequence(nums)
	}
	fmt.Println(ans)
	return ans
}

func longestIncreasingSequence(nums []int) int {
	var a = []int{}
	for _, num := range nums {
		i := upperBound(a, num)
		if i == len(a) {
			a = append(a, num)
		} else {
			a[i] = num
		}
	}
	return len(a)
}

func upperBound(nums []int, x int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r - l)>>1
		if nums[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	kIncreasing([]int{5,4,3,2,1}, 1)
	kIncreasing([]int{4,1,5,2,6,2}, 2)
	kIncreasing([]int{4,1,5,2,6,2}, 3)
	kIncreasing([]int{12,6,12,6,14,2,13,17,3,8,11,7,4,11,18,8,8,3}, 1) 	// 12
}

//[12,6,12,6,14,2,13,17,3,8,11,7,4,11,18,8,8,3]
//1
