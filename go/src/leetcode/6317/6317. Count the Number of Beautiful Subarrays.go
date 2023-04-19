package main

import "fmt"

/*
*
https://leetcode.cn/problems/count-the-number-of-beautiful-subarrays/
思路：美丽数组 <=> 所有元素异或为0
子数组num[i:j]异或和为0，等价于存在nums[:i]的异或和等于nums[:j]的异或和，i<j
*/
func beautifulSubarrays(nums []int) (tot int64) {
	cnt := map[int]int{0: 1}
	var t int
	for _, x := range nums {
		t ^= x
		if _, ok := cnt[t]; ok {
			tot += int64(cnt[t])
		}
		cnt[t]++
	}

	fmt.Println(tot)
	return
}

func main() {
	beautifulSubarrays([]int{4, 3, 1, 2, 4})
	beautifulSubarrays([]int{1, 10, 4})
}
