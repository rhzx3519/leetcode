/**
@author ZhengHao Lou
Date    2021/11/16
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/process-restricted-friend-requests/
思路：并查集
 */
func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	// 某人的敌人
	enemy := make([]map[int]int, n)
	// 某个朋友圈的所有人员
	friends := make([]map[int]int, n)
	// 某个朋友圈的所有人员的所有敌人
	friendsEnemey := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		enemy[i] = make(map[int]int)
		friends[i] = make(map[int]int)
		friendsEnemey[i] = make(map[int]int)
	}

	for _, r := range restrictions {
		u, v := r[0], r[1]
		enemy[u][v]++
		enemy[v][u]++
		friendsEnemey[u][v]++
		friendsEnemey[v][u]++
	}

	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	result := make([]bool, len(requests))
	for i, request := range requests {
		u, v := request[0], request[1]
		pu := find(parent, u)
		pv := find(parent, v)
		if pu == pv {
			result[i] = true
			continue
		}
		// 更新pu, pv的朋友圈信息
		friends[pu][u]++
		friends[pv][v]++
		for w, _ := range enemy[u] {
			friendsEnemey[pu][w]++
		}
		for w, _ := range enemy[v] {
			friendsEnemey[pv][w]++
		}

		if isjoint(friends[pu], friendsEnemey[pv]) || isjoint(friends[pv], friendsEnemey[pu]) {
			result[i] = false
			continue
		}
		// pu, pv交朋友
		union(parent, pu, pv)
		for w, _ := range friends[pv] {
			friends[pu][w]++
		}
		for w, _ := range friendsEnemey[pv] {
			friendsEnemey[pu][w]++
		}
		result[i] = true
	}

	fmt.Println(result)
	return result
}

func isjoint(friends, enemy map[int]int) bool {
	for k, _ := range friends {
		if enemy[k] != 0 {
			return true
		}
	}
	return false
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent []int, x, y int) {
	px, py := find(parent, x), find(parent, y)
	if px != py {
		parent[py] = px
	}
}

func main() {
	friendRequests(3, [][]int{{0,1}}, [][]int{{0,2},{2,1}})
	friendRequests(3, [][]int{{0,1}}, [][]int{{1,2},{0,2}})
	friendRequests(5, [][]int{{0,1},{1,2},{2,3}}, [][]int{{0,4},{1,2},{3,1},{3,4}})
}
