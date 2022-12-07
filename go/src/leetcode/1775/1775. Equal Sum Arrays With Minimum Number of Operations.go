/*
*
@author ZhengHao Lou
Date    2022/12/07
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/equal-sum-arrays-with-minimum-number-of-operations/
s1 = s2 + K
s1 - s2 = K
*/
const N = 6

func minOperations(nums1 []int, nums2 []int) (tot int) {
	n1, n2 := len(nums1), len(nums2)
	var ct1, ct2 = make([]int, N+1), make([]int, N+1)
	var s1, s2 int
	for i := range nums1 {
		s1 += nums1[i]
		ct1[nums1[i]]++
	}
	for i := range nums2 {
		s2 += nums2[i]
		ct2[nums2[i]]++
	}

	if n1 > n2*6 || n1*6 < n2 {
		return -1
	}

	var k = s1 - s2
	if k == 0 {
		return 0
	}
	if k < 0 {
		return minOperations(nums2, nums1)
	}
	var i, j = N, 1

	for ; k != 0; i, j = i-1, j+1 {
		if k < ct1[i]*(i-1) {
			tot += k / (i - 1)
			if k%(i-1) != 0 {
				tot++
			}
			return
		}
		k -= ct1[i] * (i - 1)
		tot += ct1[i]

		if k < ct2[j]*(N-j) {
			tot += k / (N - j)
			if k%(N-j) != 0 {
				tot++
			}
			return
		}
		k -= ct2[j] * (N - j)
		tot += ct2[j]
	}

	return
}

func main() {
	fmt.Println(minOperations([]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2}))
	//minOperations([]int{1, 1, 1, 1, 1, 1, 1}, []int{6})
	fmt.Println(minOperations([]int{6, 6}, []int{1}))
}
