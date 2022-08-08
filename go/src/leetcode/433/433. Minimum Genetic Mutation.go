/**
@author ZhengHao Lou
Date    2022/05/07
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/minimum-genetic-mutation/
*/
func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	has := make(map[string]bool)
	for _, b := range bank {
		has[b] = true
	}
	if !has[end] {
		return -1
	}
	
	que := []string{start}
	for step := 0; len(que) != 0; step++ {
		for i := len(que) - 1; i >= 0; i-- {
			s := que[0]
			que = que[1:]
			if s == end {
				return step
			}
			var keys []string
			for t := range has {
				var x int
				for i := range s {
					if s[i] != t[i] {
						x++
					}
				}
				if x == 1 {
					keys = append(keys, t)
					que = append(que, t)
				}
			}
			for _, t := range keys {
				delete(has, t)
			}
		}
	}
	
	return -1
}

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGATT", "AACCGATA", "AAACGATA", "AAACGGTA"}))
}
