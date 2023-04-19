package main

import "fmt"

/**
https://leetcode-cn.com/problems/all-paths-from-source-to-target/
思路：dfs
 */
func allPathsSourceTarget(graph [][]int) [][]int {
	ans := [][]int{}
	n := len(graph)

	var dfs func(int, []int)
	dfs = func(i int, path []int) {
		if i == n-1 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for _, ni := range graph[i] {
			path = append(path, ni)
			dfs(ni, path)
			path = path[:len(path) - 1]
		}
	}
	dfs(0, []int{0})
	return ans
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1,2},{3},{3},{}}))
}
