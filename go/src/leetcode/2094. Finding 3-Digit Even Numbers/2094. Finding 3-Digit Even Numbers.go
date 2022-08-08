/**
@author ZhengHao Lou
Date    2021/12/06
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/finding-3-digit-even-numbers/
 */
func findEvenNumbers(digits []int) []int {
	var mapper = make(map[int]int)

	for i := range digits {
		if digits[i] == 0 {
			continue
		}
		for j := range digits {
			if i == j || digits[j]&1 == 1 {
				continue
			}
			for k := range digits {
				if i == k || j == k {
					continue
				}
				x := 100*digits[i] + 10*digits[k] + digits[j]
				mapper[x]++
			}
		}
	}
	var ans = []int{}
	for k, _ := range mapper {
		ans = append(ans, k)
	}
	sort.Ints(ans)

	fmt.Println(ans)
	return ans
}

func main() {
	//findEvenNumbers([]int{2,1,3,0})
	findEvenNumbers([]int{2,2,8,8,2})
	findEvenNumbers([]int{3,7,5})
}
