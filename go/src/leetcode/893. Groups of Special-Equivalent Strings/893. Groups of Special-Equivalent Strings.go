/**
@author ZhengHao Lou
Date    2021/11/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/groups-of-special-equivalent-strings/
 */
const N = 26

type word struct {
	even [N]int
	odd [N]int
}

func numSpecialEquivGroups(words []string) int {
	counter := make(map[word]int)
	for _, w := range words {
		even, odd := extractTo(w)
		counter[word{even: even, odd: odd}]++
	}

	return len(counter)
}

func extractTo(w string) ([N]int, [N]int) {
	even, odd := [N]int{}, [N]int{}
	for i, j := 0, 1; i < len(w) || j < len(w); i, j = i+2, j+2 {
		if i < len(w) {
			even[int(w[i] - 'a')]++
		}
		if j < len(w) {
			odd[int(w[j] - 'a')]++
		}
	}
	return even, odd
}

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"abcd","cdab","cbad","xyzz","zzxy","zzyx"}))
	fmt.Println(numSpecialEquivGroups([]string{"abc","acb","bac","bca","cab","cba"}))
}
