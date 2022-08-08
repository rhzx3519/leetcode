/**
@author ZhengHao Lou
Date    2021/11/10
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/distant-barcodes/
思路：按数量从大到小排列，先设置偶数位，再设置奇数位
 */
func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
	mapper := map[int]int{}
	for _, b := range barcodes {
		mapper[b]++
	}

	bs := []int{}
	for i, _ := range mapper {
		bs = append(bs, i)
	}
	sort.Slice(bs, func(i, j int) bool {
		return mapper[bs[i]] > mapper[bs[j]]
	})

	var j, c int
	var ans = make([]int, n)
	for i := 0; i < n; i += 2 {
		ans[i] = bs[j]
		c++
		if c == mapper[bs[j]] {
			j++
			c = 0
		}
	}

	for i := 1; i < n; i += 2 {
		ans[i] = bs[j]
		c++
		if c == mapper[bs[j]] {
			j++
			c = 0
		}
	}

	return ans
}


func main() {
	fmt.Println(rearrangeBarcodes([]int{1,1,1,2,2,2}))
	fmt.Println(rearrangeBarcodes([]int{1,1,1,1,2,2,3,3}))
}
