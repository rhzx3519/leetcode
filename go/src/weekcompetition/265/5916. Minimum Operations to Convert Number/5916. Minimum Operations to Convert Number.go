/**
@author ZhengHao Lou
Date    2021/11/01
*/
package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/minimum-operations-to-convert-number/
思路：BFS + 剪枝
记录访问过的节点，再次访问时直接跳过
 */
const N int = 1000
func minimumOperations(nums []int, start int, goal int) int {
	vis := make(map[int]bool)
	vis[start] = true
	var steps = 1
	que := []int{start}
	for len(que) != 0 {
		l := len(que)
		for tmp := 0; tmp < l; tmp++ {
			x := que[0]
			que = que[1:]
			for _, num := range nums {
				nexts := []int{x + num, x - num, x ^ num}
				for _, next := range nexts {
					if next == goal {
						return steps
					}
					if next >= 0 && next <= N && !vis[next] {
						vis[next] = true
						que = append(que, next)
					}
				}
			}
		}
		steps++
	}

	return -1
}

func main() {
	fmt.Println(minimumOperations([]int{1,3}, 6, 4))
	fmt.Println(minimumOperations([]int{2,4,12}, 2, 12))
	fmt.Println(minimumOperations([]int{3,5,7}, 0, -4))
	fmt.Println(minimumOperations([]int{2,8,16}, 0, 1))
	fmt.Println(minimumOperations([]int{1}, 0, 3))
}
