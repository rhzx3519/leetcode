/**
@author ZhengHao Lou
Date    2022/03/28
*/
package main

/**
https://leetcode-cn.com/problems/find-the-difference-of-two-arrays/
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
	var foo = func(nums1 []int, nums2 []int) []int {
		var arr []int
		counter := make(map[int]int)
		for _, num := range nums1 {
			counter[num]++
		}
		for _, num := range nums2 {
			counter[num] = 0
		}

		for num, c := range counter {
			if c > 0 {
				arr = append(arr, num)
			}
		}
		return arr
	}

	var ans = make([][]int, 2)
	ans[0] = foo(nums1, nums2)
	ans[1] = foo(nums2, nums1)

	return ans
}

//[[-80,-15,-81,-61,14,-35,-10,-28,63,-45],[-1,-40,-44,41,10,-43,69,10,2]]
//[[-81,-35,-10,-28,-61,-45,-15,14,-80,63],[-1,2,69,-40,41,10,-43,-44]]
