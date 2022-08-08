/**
@author ZhengHao Lou
Date    2022/02/05
*/
package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	var lower, upper, equal []int
	for _, num := range nums {
		if num < pivot {
			lower = append(lower, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			upper = append(upper, num)
		}
	}
	var ans = make([]int, len(nums))
	copy(ans, lower)
	copy(ans[len(lower):], equal)
	copy(ans[len(lower)+len(equal):], upper)

	fmt.Println(ans)
	return ans
}

func main() {
	pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10)
	pivotArray([]int{-3, 4, 3, 2}, 2)
}
