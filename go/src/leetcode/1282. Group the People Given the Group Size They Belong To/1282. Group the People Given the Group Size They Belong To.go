/**
@author ZhengHao Lou
Date    2021/12/02
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/group-the-people-given-the-group-size-they-belong-to/
 */
func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)
	for i, size := range groupSizes {
		if _, ok := groups[size]; !ok {
			groups[size] = []int{}
		}
		groups[size] = append(groups[size], i)
	}

	var users = [][]int{}
	for sz, arr := range groups {
		if len(arr)	> sz {
			for i := 0; i < len(arr); i += sz {
				users = append(users, append([]int{}, arr[i:i+sz]...))
			}
		} else {
			users = append(users, arr)
		}
	}

	return users
}

func main() {
	fmt.Println(groupThePeople([]int{3,3,3,3,3,1,3}))
	fmt.Println(groupThePeople([]int{2,1,3,3,3,2}))
}