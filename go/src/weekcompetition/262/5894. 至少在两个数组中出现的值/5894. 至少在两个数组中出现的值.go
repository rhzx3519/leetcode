/**
@author ZhengHao Lou
Date    2021/10/10
*/
package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	mapper := map[int]int{}
	tmp := map[int]int{}
	for _, num := range nums1 {
		if tmp[num] == 0 {
			tmp[num] = 1
			mapper[num]++
		}
	}

	tmp = map[int]int{}
	for _, num := range nums2 {
		if tmp[num] == 0 {
			tmp[num] = 1
			mapper[num]++
		}
	}

	tmp = map[int]int{}
	for _, num := range nums3 {
		if tmp[num] == 0 {
			tmp[num] = 1
			mapper[num]++
		}
	}

	ans := []int{}
	for num, count := range mapper {
		if count >= 2 {
			ans = append(ans, num)
		}
	}
	return ans
}
