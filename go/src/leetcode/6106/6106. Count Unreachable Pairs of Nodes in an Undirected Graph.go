/**
@author ZhengHao Lou
Date    2022/06/27
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
*/
func countPairs(n int, edges [][]int) int64 {
	parent := make([]int, n)
	counter := make(map[int]int)
	for i := 0; i < n; i++ {
		parent[i] = i
		counter[i] = 1
	}

	//fmt.Println(parent)

	var find func(x int) int
	var union func(x, y int)

	find = func(x int) int {
		if parent[x] != x {
			return find(parent[x])
		}
		return x
	}

	union = func(x, y int) {
		px, py := find(x), find(y)
		if px != py {
			parent[py] = px
			counter[px] += counter[py]
			delete(counter, py)
		}
	}
	//fmt.Println(counter)

	for _, e := range edges {
		union(e[0], e[1])
	}

	var nums = []int{}
	var s int
	for _, v := range counter {
		s += v
		nums = append(nums, v)
	}

	var ans int64
	for i := range nums {
		s -= nums[i]
		ans += int64(nums[i] * s)
	}

	fmt.Println(ans)
	return ans
}

func main() {
	countPairs(7, [][]int{{0, 2}, {5, 0}, {2, 4}, {1, 6}, {5, 4}})
	countPairs(100000, [][]int{})
}
