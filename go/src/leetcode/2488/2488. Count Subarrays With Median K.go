/*
*
@author ZhengHao Lou
Date    2022/11/28
*/
package main

/*
*
https://leetcode.cn/problems/count-subarrays-with-median-k/
la + ra = lb + rb
la - lb = rb - ra
OR
la + ra + 1 = lb + rb
la - lb + 1 = rb - ra
*/
func countSubarrays(nums []int, k int) (tot int) {
	var ik int
	for i := range nums {
		if nums[i] == k {
			ik = i
			break
		}
	}
	n := len(nums)
	var left = map[int]int{0: 1}
	var a, b int
	for i := ik - 1; i >= 0; i-- {
		if nums[i] < k {
			a++
		} else {
			b++
		}
		left[a-b]++
	}
	a, b = 0, 0
	for i := ik; i < n; i++ {
		if nums[i] < k {
			a++
		} else if nums[i] > k {
			b++
		}
		tot += left[b-a]

		tot += left[b-a-1]
	}

	//fmt.Println(tot, left)
	return
}

func main() {
	countSubarrays([]int{3, 2, 1, 4, 5}, 4)
	countSubarrays([]int{2, 3, 1}, 3)
}
