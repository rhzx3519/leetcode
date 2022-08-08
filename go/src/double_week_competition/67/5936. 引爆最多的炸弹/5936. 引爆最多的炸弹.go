/**
@author ZhengHao Lou
Date    2021/12/11
*/
package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/contest/biweekly-contest-67/problems/detonate-the-maximum-bombs/
思路：建图+DFS/BFS
1. B的圆心在A的半径内 ，A可以引爆B，表示存在A指向B的边，按照这个原则建图
2. 从每个节点开始遍历，求出最大引爆数量
 */
func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && isIntersect(bombs[i], bombs[j]){
				if _, ok := graph[i]; !ok {
					graph[i] = []int{}
				}
				graph[i] = append(graph[i], j)
			}
		}
	}
	// 求连通域
	vis := make([]bool, n)
	var dfs func(idx int) int
	dfs = func(idx int) int {
		var c = 1
		vis[idx] = true
		for _, ni := range graph[idx] {
			if !vis[ni] {
				c += dfs(ni)
			}
		}
		return c
	}

	var  ans int
	for i := 0; i < n; i++ {
		vis = make([]bool, n)
		c := dfs(i)
		if ans < c {
			ans = c
		}
	}
	fmt.Println(ans)
	return ans
}

func isIntersect(a, b []int) bool {
	t := (b[0] - a[0])*(b[0] - a[0]) + (b[1] - a[1])*(b[1] - a[1])
	return t <= a[2]*a[2]
}


//[[855,82,158],[17,719,430],[90,756,164],[376,17,340],[691,636,152],[565,776,5],[464,154,271],[53,361,162]
//,[278,609,82],[202,927,219],[542,865,377],[330,402,270],[720,199,10],[986,697,443],[471,296,69],[393,81,404],[127,405,177]]
func main() {
	//fmt.Println(isIntersect([]int{691,636,152}, []int{565,776,5}))
	maximumDetonation([][]int{{855,82,158},{17,719,430},{90,756,164},{376,17,340},{691,636,152},{565,776,5},{464,154,271},{53,361,162},{278,609,82},{202,927,219},{542,865,377},{330,402,270},{720,199,10},{986,697,443},{471,296,69},{393,81,404},{127,405,177}})
	maximumDetonation([][]int{{2,1,3},{6,1,4}})
	maximumDetonation([][]int{{1,1,5},{10,10,5}})
	maximumDetonation([][]int{{1,2,3},{2,3,1},{3,4,2},{4,5,3},{5,6,4}})
	maximumDetonation([][]int{{54,95,4},{99,46,3},{29,21,3},{96,72,8},{49,43,3},{11,20,3},{2,57,1},{69,51,7},{97,1,10},{85,45,2},{38,47,1},{83,75,3},{65,59,3},{33,4,1},{32,10,2},{20,97,8},{35,37,3}})
	maximumDetonation([][]int{{7,26,7},{7,18,4},{3,10,7},{17,50,1},{3,25,10},{85,23,8},{80,50,1},{58,74,1},{38,39,7},{50,51,8},{31,99,3},{53,6,5},{59,27,10},{87,78,9},{68,58,3}})
}