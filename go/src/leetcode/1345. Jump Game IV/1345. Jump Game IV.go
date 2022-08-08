/**
@author ZhengHao Lou
Date    2022/01/21
*/
package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/jump-game-iv/
思路：简单BFS
两个剪枝技巧：
1. 使用vis记录访问的节点
2. bfs先访问的节点肯定是更近的，所以可以删除相同值对应的counter
 */
func minJumps(arr []int) int {
	n := len(arr)
	counter := make(map[int][]int)
	for i := n-1; i >= 0; i-- {
		if _, ok := counter[arr[i]]; !ok {
			counter[arr[i]] = []int{}
		}
		counter[arr[i]] = append(counter[arr[i]], i)
	}

	fmt.Println(counter)
	vis := make([]bool, n)
	vis[0] = true
	var step int
	que := []int{0}
	for len(que) != 0 {
		l := len(que)
		for i := 0; i < l; i++ {
			x := que[0]
			que = que[1:]
			if x == n-1 {
				return step
			}
			for j := len(counter[arr[x]])-1; j >= 0; j-- {
				y := counter[arr[x]][j]
				if vis[y] {
					continue
				}
				vis[y] = true
				que = append(que, y)
			}
			delete(counter, arr[x])	// 剪枝
			if x + 1 < n && !vis[x+1] {
				vis[x+1] = true
				que = append(que, x+1)
			}
			if x - 1 >= 0 && !vis[x-1] {
				vis[x-1] = true
				que = append(que, x-1)
			}
		}
		step++
	}

	return 0
}


func main() {
	fmt.Println(minJumps([]int{100,-23,-23,404,100,23,23,23,3,404}))
	fmt.Println(minJumps([]int{7}))
	fmt.Println(minJumps([]int{7,6,9,6,9,6,9,7}))
	fmt.Println(minJumps([]int{6,1,9}))
}