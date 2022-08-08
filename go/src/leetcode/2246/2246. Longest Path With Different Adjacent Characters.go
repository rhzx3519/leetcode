/**
@author ZhengHao Lou
Date    2022/04/20
*/
package main

/**
https://leetcode-cn.com/problems/longest-path-with-different-adjacent-characters/
思路：dfs
*/
func longestPath(parent []int, s string) int {
    n := len(parent)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        g[parent[i]] = append(g[parent[i]], i)
    }

    var ans int
    var dfs func(i int) int // 返回从i出发的一条合法最长路径长度
    dfs = func(i int) (mx int) {
        for _, ni := range g[i] {
            l := dfs(ni) + 1 // 从ni出发的最长合法路径长度
            if s[i] != s[ni] {
                ans = max(ans, mx+l) // 以i为连接点的i的子树中两条最长路径之和
                mx = max(mx, l)
            }
        }
        return
    }

    dfs(0)

    return ans + 1 // +1 表示加上连接点
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    //longestPath([]int{-1, 0, 0, 1, 1, 2}, "abacbe")
    //longestPath([]int{-1, 0, 0, 0}, "aabc")
    longestPath([]int{-1, 0, 1}, "aab")
}
