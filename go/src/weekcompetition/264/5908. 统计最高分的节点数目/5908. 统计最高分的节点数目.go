/**
@author ZhengHao Lou
Date    2021/10/24
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-nodes-with-the-highest-score/
 */
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	tree := make([][]int, n)
	for i := range tree {
		tree[i] = []int{}
	}
	for i := range parents {
		if parents[i] == -1 {
			continue
		}
		tree[parents[i]] = append(tree[parents[i]], i)
	}

	fmt.Println(tree)

	// 统计以id为根结点的二叉树节点数目
	mem := make(map[int]int)
	var dfs func(id int) int
	dfs = func(id int) int {
		if mem[id] != 0 {
			return mem[id]
		}
		var count = 1
		for _, child := range tree[id] {
			count += dfs(child)
		}
		mem[id] = count
		return count
	}
	dfs(0)

	fmt.Println(mem)
	var maxScore int64
	var count int
	for i := 0; i < n; i++ {
		var x int64 = 1
		if n - mem[i] != 0 {
			x *= int64(n - mem[i])
		}
		for _, child := range tree[i] {
			if mem[child] != 0 {
				x *= int64(mem[child])
			}
		}
		if x > maxScore {
			maxScore = x
			count = 1
		} else if x == maxScore {
			count++
		}
	}

	fmt.Println(count)
	return count
}

func main() {
	//countHighestScoreNodes([]int{-1,2,0,2,0}) // 3
	countHighestScoreNodes([]int{-1,2,0})	// 2
}
