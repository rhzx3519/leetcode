/**
@author ZhengHao Lou
Date    2022/01/28
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/the-number-of-weak-characters-in-the-game/
思路：按照攻击力递减排序，攻击力相同则按防御力递增排序
从前往后遍历properties，同时记录当前最大防御力，如果存在当前角色防御力弱于最大防御力，则弱角色+1；否则更新最大防御力
*/
func numberOfWeakCharacters(properties [][]int) int {
	var count int
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] != properties[j][0] {
			return properties[i][0] > properties[j][0]
		}
		return properties[i][1] < properties[j][1]
	})
	var maxDef int
	for i := range properties {
		if properties[i][1] < maxDef {
			count++
		} else {
			maxDef = properties[i][1]
		}

	}
	fmt.Println(count)
	return count
}

func main() {
	numberOfWeakCharacters([][]int{{5, 5}, {6, 3}, {3, 6}})
	numberOfWeakCharacters([][]int{{2, 2}, {3, 3}})
	numberOfWeakCharacters([][]int{{1, 5}, {10, 4}, {4, 3}})
}
