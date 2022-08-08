/**
@author ZhengHao Lou
Date    2021/12/27
*/
package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/friends-of-appropriate-ages/
1. ages[y] > 0.5*ages[x] + 7
2. ages[y] <= ages[x]
 */
const N = 125
func numFriendRequests(ages []int) int {
	var req int
	slots := make([]int, N)

	for _, age := range ages {
		slots[age]++
	}
	for i := 1; i <= 120; i++ {
		if slots[i] == 0 {
			continue
		}
		for j := i>>1+8; j <= i; j++ {
			if slots[j] == 0 {
				continue
			}
			if i < 100 && j > 100 {
				break
			}
			if i == j {
				req += (slots[i] - 1)*slots[i]
			} else {
				req += slots[i] * slots[j]
			}
			//fmt.Println(i, j)
		}
	}
	return req
}


func main() {
	fmt.Println(numFriendRequests([]int{16,16}))
	fmt.Println(numFriendRequests([]int{16,17,18}))
	fmt.Println(numFriendRequests([]int{20,30,100,110,120}))
}