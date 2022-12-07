/**
@author ZhengHao Lou
Date    2022/09/26
*/
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/number-of-good-paths/
按节点值从小到大考虑，同时用并查集合并时，总是从节点值小的点往节点值大的点合并，这样可以保证连通块的代表元的节点值是最大的。
对于节点 x 及其邻居 y，如果 y 所处的连通分量的最大节点值不超过 vals[x]，那么可以把 y 所处的连通块合并到 x 所处的连通块中。
如果此时这两个连通块的最大节点值相同，那么可以根据乘法原理，把这两个连通块内的等于最大节点值的节点个数相乘，加到答案中。

作者：endlesscheng
链接：https://leetcode.cn/problems/number-of-good-paths/solution/bing-cha-ji-by-endlesscheng-tbz8/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func numberOfGoodPaths(vals []int, edges [][]int) (c int) {
	n := len(vals)
	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	nums := make([]int, n)
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
		parent[i] = i
		size[i] = 1
	}
	sort.Slice(nums, func(i, j int) bool {
		return vals[nums[i]] < vals[nums[j]]
	})

	var find func(x int) int
	find = func(x int) int {
		if x != parent[x] {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	c = n
	for _, i := range nums {
		pi := find(i)
		for _, j := range adj[i] {
			pj := find(j)
			if pi == pj || vals[j] > vals[i] {
				continue
			}

			if vals[pi] == vals[pj] {
				c += size[pi] * size[pj]
				size[pi] += size[pj]
			}

			parent[pj] = pi
		}
	}

	fmt.Println(parent, size, c)
	return
}

func main() {
	numberOfGoodPaths([]int{1, 3, 2, 1, 3}, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}})
	numberOfGoodPaths([]int{1, 1, 2, 2, 3}, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}})
	numberOfGoodPaths([]int{1}, [][]int{})
}
