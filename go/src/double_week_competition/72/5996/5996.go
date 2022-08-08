/**
@author ZhengHao Lou
Date    2022/02/19
*/
package main

func countPairs(nums []int, k int) int {
	var c int
	counter := make(map[int][]int)
	for i, num := range nums {
		for _, j := range counter[num] {
			if (i*j)%k == 0 {
				c++
			}
		}
		counter[num] = append(counter[num], i)
	}

	return c
}
