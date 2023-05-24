package main

/*
*
https://leetcode.cn/problems/frog-position-after-t-seconds/description/
*/
func frogPosition(n int, edges [][]int, t int, target int) float64 {
    g := make([][]int, n+1)
    g[1] = []int{0} // 给结点1添加邻居0，避免特殊判断
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    var dfs func(x, fa, leftT int) int
    dfs = func(x, fa, leftT int) int {
        if leftT == 0 {
            if x == target {
                return 1
            }
            return 0
        }
        if x == target {
            if len(g[x]) == 1 {
                return 1
            }
            return 0
        }
        for _, nx := range g[x] {
            if nx != fa {
                prod := dfs(nx, x, leftT-1)
                if prod != 0 {
                    return prod * (len(g[x]) - 1) // 乘上儿子的个数
                }
            }
        }
        return 0;
    }

    prod := dfs(1, 0, t)
    if prod == 0 {
        return 0
    }
    return 1 / float64(prod)
}


