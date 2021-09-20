package main

import (
	"fmt"
	"sort"
)

func findOriginalArray(changed []int) []int {
	n := len(changed)
	mapper := make(map[int]int)
	for _, num := range changed {
		mapper[num]++
	}

	sort.Ints(changed)

	var ans = []int{}
	for i := n - 1; i >= 0; i-- {
		num := changed[i]
		if _, ok := mapper[num]; !ok {
			continue
		}
		if num&1 == 0 {
			if _, ok := mapper[num>>1]; ok {
				ans = append(ans, num>>1)
				mapper[num>>1]--
				mapper[num]--
				if mapper[num] == 0 {
					delete(mapper, num)
				}
				if mapper[num>>1] == 0 {
					delete(mapper, num>>1)
				}
			}
		}
	}

	if len(mapper) != 0 {
		return []int{}
	}
	return ans
}

func main() {
	fmt.Println(findOriginalArray([]int{1,3,4,2,6,8}))
	fmt.Println(findOriginalArray([]int{6,3,0,1}))
	fmt.Println(findOriginalArray([]int{1}))
}
