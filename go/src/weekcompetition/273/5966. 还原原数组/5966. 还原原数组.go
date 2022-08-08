/**
@author ZhengHao Lou
Date    2021/12/26
*/
package main

import (
	"fmt"
	"sort"
)

/**
	lower[i] + k = higher[i] - k
high[i] - lower[i] = 2*k
 */
func recoverArray(nums []int) []int {
	sort.Ints(nums)
	n := len(nums) >> 1
	var ans = make([]int, n)


	for i := 1; i < n<<1; i++ {
		if nums[i] - nums[0] == 0 || (nums[i] - nums[0]) & 1 != 0 {
			continue
		}
		counter := make(map[int]int)
		for _, num := range nums {
			counter[num]++
		}
		var tmp = []int{}
		k := (nums[i] - nums[0])>>1
		for i := 0; i < n<<1; i++ {
			if counter[nums[i]] == 0 {
				continue
			}
			if counter[nums[i] + k<<1] == 0 {
				break
			}
			counter[nums[i]]--
			counter[nums[i] + k<<1]--
			tmp = append(tmp, nums[i] + k)
		}

		if len(tmp) == n {
			fmt.Println(tmp, k)
			return tmp
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	recoverArray([]int{2,10,6,4,8,12})
	recoverArray([]int{1,1,3,3})
	recoverArray([]int{5,435})
}
