/**
@author ZhengHao Lou
Date    2022/05/03
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
https://leetcode-cn.com/problems/k-divisible-elements-subarrays/
*/
func countDistinct(nums []int, k int, p int) int {
	var counter = make(map[string]int)
	var strs []string
	for i := range nums {
		strs = append(strs, strconv.Itoa(nums[i]))
	}
	
	var index = []int{}
	var j int
	for i := range nums {
		if nums[i]%p == 0 {
			index = append(index, i)
		}
		if len(index) > k {
			j = index[0] + 1
			index = index[1:]
		}
		
		for k := j; k <= i; k++ {
			counter[strings.Join(strs[k:i+1], ",")]++
		}
	}
	
	fmt.Println(len(counter))
	fmt.Println(counter)
	return len(counter)
}

func main() {
	countDistinct([]int{2, 3, 2, 3, 2, 2}, 2, 2) // 10
	//countDistinct([]int{2, 3, 3, 2, 2}, 2, 2)
	//countDistinct([]int{1, 2, 3, 4}, 4, 1)
}
