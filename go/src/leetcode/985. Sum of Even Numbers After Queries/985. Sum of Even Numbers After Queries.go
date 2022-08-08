/**
@author ZhengHao Lou
Date    2021/11/25
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/sum-of-even-numbers-after-queries/
 */
func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	var even int
	for _, num := range nums {
		if num&1 == 0 {
			even += num
		}
	}
	var ans = []int{}
	for _, q := range queries {
		val, index := q[0], q[1]
		d := nums[index]
		nums[index] += val
		if nums[index] & 1 == 0 {
			if d & 1 == 0 {
				even += val
			} else {
				even += nums[index]
			}
		} else {
			if d & 1 == 0 {
				even -= d
			}
		}
		ans = append(ans, even)
	}

	return ans
}

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1,2,3,4}, [][]int{{1,0},{-3,1},{-4,0},{2,3}}))
}
