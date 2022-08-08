/**
@author ZhengHao Lou
Date    2022/05/24
*/
package main

/**
https://leetcode.cn/problems/maximum-score-of-a-node-sequence/
*/
func maximumScore(scores []int, edges [][]int) int {
	adj := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	
	//  返回节点idx，左边长度为l，右边为r的左右节点最大长度
	var dfs func(idx, l, r int) (int, int)
	dfs = func(idx, l, r int) (int, int) {
	
	}
	
}
