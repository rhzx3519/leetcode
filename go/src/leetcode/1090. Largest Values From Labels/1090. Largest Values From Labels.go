/**
@author ZhengHao Lou
Date    2021/11/09
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/largest-values-from-labels/
 */
type item struct {
	id, value, label int
}
func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	n := len(values)
	items := []item{}
	for i := 0; i < n; i++ {
		items = append(items, item{id: i, value: values[i], label: labels[i]})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].value < items[j].value
	})

	var score int
	var subsets = []int{}
	limit := make(map[int]int)
	for len(items) != 0 && len(subsets) < num_wanted {
		it := items[len(items) - 1]
		if limit[it.label] < use_limit {
			subsets = append(subsets, it.id)
			score += it.value
			limit[it.label]++
			//fmt.Println(it.id, it.label)
		}
		items = items[:len(items) - 1]
	}
	return score
}


func main() {
	fmt.Println(largestValsFromLabels([]int{5,4,3,2,1}, []int{1,1,2,2,3}, 3, 1))
	fmt.Println(largestValsFromLabels([]int{5,4,3,2,1}, []int{1,3,3,3,2}, 3, 2))
	fmt.Println(largestValsFromLabels([]int{9,8,8,7,6}, []int{0,0,0,1,1}, 3, 1))
	fmt.Println(largestValsFromLabels([]int{9,8,8,7,6}, []int{0,0,0,1,1}, 3, 2))
}






