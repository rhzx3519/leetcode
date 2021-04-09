/**
思路：贪心法
 */
package main

import "fmt"

func canChoose(groups [][]int, nums []int) bool {
	m, n := len(groups), len(nums)
	count := 0
	for i := 0; i < n && count < m; {
		if groups[count][0] == nums[i] {
			if equal(groups[count], nums, i) {
				i += len(groups[count])
				count++
				continue
			}
		}

		i++
	}

	return count == m
}

func equal(g, nums []int, i int) bool {
	k := 0
	for ; k < len(g) && i+k < len(nums); k++ {
		if g[k] != nums[i+k] {
			return false
		}
	}
	return k == len(g)
}

func main() {
	fmt.Println(canChoose([][]int{{1,-1,-1}, {3,-2,0}}, []int{1,-1,0,1,-1,-1,3,-2,0}))
	fmt.Println(canChoose([][]int{{10,-2}, {1,2,3,4}}, []int{1,2,3,4,10,-2}))
	fmt.Println(canChoose([][]int{{1,2,3}, {3,4}}, []int{7,7,1,2,3,4,7,7}))
}
