package main

/**
https://leetcode.cn/problems/count-pairs-of-connectable-servers-in-a-weighted-tree-network/?envType=daily-question&envId=2024-06-04
*/
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
    type edge struct {
        to, wt int
    }
    adj := make([][]edge, len(edges))
    for _, e := range edges {
        a, b, wt := e[0], e[1], e[2]
        adj[a] = append(adj[a], edge{b, wt})
        adj[b] = append(adj[b], edge{a, wt})
    }

    var ans = make([]int, len(edges))
    for i := range adj {
        var cnt int
        var dfs func(int, int, int)
        dfs = func(i, fi, sum int) {
            if sum%signalSpeed == 0 {
                cnt++
            }
            for _, e := range adj[i] {
                if e.to != fi {
                    dfs(e.to, i, sum+e.wt)
                }
            }
        }

        var sum int
        for _, e := range adj[i] {
            cnt = 0
            dfs(e.to, i, e.wt)
            ans[i] += cnt * sum
            sum += cnt
        }
    }
    return ans
}
