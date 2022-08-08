/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/sort-integers-by-the-power-value/
 */
func getKth(lo int, hi int, k int) int {
	weight := map[int]int{}
	nums := []int{}
	for i := lo; i <= hi; i++ {
		weight[i] = getWeight(i)
		nums = append(nums, i)
	}
	sort.Slice(nums, func(i, j int) bool {
		if weight[nums[i]] != weight[nums[j]] {
			return weight[nums[i]] < weight[nums[j]]
		}
		return nums[i] < nums[j]
	})
	//fmt.Println(weight)
	//fmt.Println(nums)
	fmt.Println(nums[k - 1])
	return nums[k - 1]
}

func getWeight(x int) int {
	var w int
	for  x != 1 {
		if x & 1 == 1 {
			x = x<<1 + x + 1
		} else {
			x >>= 1
		}
		w++
	}
	return w
}

func main() {
	//getKth(12, 15, 2)
	//getKth(1, 1, 1)
	//getKth(7, 11, 4)
	//getKth(10, 20, 5)
	getKth(1, 1000, 777)
}