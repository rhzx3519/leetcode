/**
@author ZhengHao Lou
Date    2022/10/30
*/
package main

import (
	"fmt"
)

/**
https://leetcode.cn/problems/destroy-sequential-targets/
思路：x + c*space = y,  (y-x)%space == 0, y%space == x%space
*/
func destroyTargets(nums []int, space int) (ans int) {
	mod := make(map[int][]int)
	for i := range nums {
		mod[nums[i]%space] = append(mod[nums[i]%space], nums[i])
	}
	var mx int
	for _, g := range mod {
		var mi int = 1e9 + 7
		for _, i := range g {
			if i < mi {
				mi = i
			}
		}
		if len(g) > mx || (len(g) == mx && mi < ans) {
			mx = len(g)
			ans = mi
		}

	}
	return ans

}

func main() {
	fmt.Println(destroyTargets([]int{3, 7, 8, 1, 1, 5}, 2))
	fmt.Println(destroyTargets([]int{1, 3, 5, 2, 4, 6}, 2))
	fmt.Println(destroyTargets([]int{6, 2, 5}, 100))
	fmt.Println(destroyTargets([]int{1, 5, 3, 2, 2}, 10000)) // 2
}
