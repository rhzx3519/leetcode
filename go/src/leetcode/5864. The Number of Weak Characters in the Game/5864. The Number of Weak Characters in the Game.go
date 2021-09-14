package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/the-number-of-weak-characters-in-the-game/
思路：将角色按照攻击力从大到小排列，如果攻击力相同，则按照防御力从小到大排列

 */
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		p1, p2 := properties[i], properties[j]
		if p1[0] != p2[0] {
			return p1[0] > p2[0]
		} else {
			return p1[1] < p2[1]
		}
	})
	var count int
	var maxDef int	// 记录i之前的最大防御力
	for _, p := range properties {
		if p[1] < maxDef {	// p之前存在防御力、攻击力大于p的角色
			count++
		} else {
			maxDef = p[1]
		}
	}

	return count
}

func main() {
	fmt.Println(numberOfWeakCharacters([][]int{{5,5},{6,3},{3,6}}))
	fmt.Println(numberOfWeakCharacters([][]int{{2,2},{3,3}}))
	fmt.Println(numberOfWeakCharacters([][]int{{2,2},{3,3}}))
	fmt.Println(numberOfWeakCharacters([][]int{{1,5},{10,4},{4,3}}))
	fmt.Println(numberOfWeakCharacters([][]int{{1,1},{2,1},{2,2},{1,2}}))
	fmt.Println(numberOfWeakCharacters([][]int{{7,7},{1,2},{9,7},{7,3},{3,10},{9,8},{8,10},{4,3},{1,5},{1,5}}))
}
