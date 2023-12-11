package main

import (
    "fmt"
    "sort"
)

/**
https://leetcode.cn/problems/path-with-minimum-effort/?envType=daily-question&envId=2023-12-11
思路：并查集，将每个格子看作是图的一个节点，相邻两个元素之间的高度差看作边的权重，
按照边的权重从小到大排序，依次添加边，直到左上角和右下角连通为止
*/
var df = []int{0, 1, 0}

type DisjoinSet struct {
    n      int
    parent []int
    size   []int
}

func (d DisjoinSet) Find(x int) int {
    if d.parent[x] != x {
        d.parent[x] = d.Find(d.parent[x])
    }
    return d.parent[x]
}

func (d DisjoinSet) Union(x, y int) bool {
    px, py := d.Find(x), d.Find(y)
    if px != py {
        if d.size[px] > d.size[py] {
            d.parent[py] = px
            d.size[px] += d.size[py]
        } else {
            d.parent[px] = py
            d.size[py] += d.size[px]
        }
        return true
    }
    return false
}

func (d DisjoinSet) Connected(a, b int) bool {
    return d.Find(a) == d.Find(b)
}

func (d DisjoinSet) Size(x int) int {
    return d.size[d.Find(x)]
}

func NewDisjointSet(n int) DisjoinSet {
    parent := make([]int, n)
    sz := make([]int, n)
    for i := range parent {
        parent[i] = i
        sz[i] = 1
    }
    return DisjoinSet{
        n:      n,
        parent: parent,
        size:   sz,
    }
}

func minimumEffortPath(heights [][]int) int {
    m, n := len(heights), len(heights[0])
    type edge struct {
        w, a, b int
    }
    var edges []edge
    for i := range heights {
        for j := range heights[i] {
            for k := 0; k < 2; k++ {
                ni, nj := i+df[k], j+df[k+1]
                if ni >= 0 && ni < m && nj >= 0 && nj < n {
                    edges = append(edges, edge{abs(heights[i][j] - heights[ni][nj]), i*n + j, ni*n + nj})
                }
            }
        }
    }
    sort.Slice(edges, func(i, j int) bool {
        return edges[i].w < edges[j].w
    })

    d := NewDisjointSet(m * n)
    for _, e := range edges {
        d.Union(e.a, e.b)
        if d.Connected(0, m*n-1) {
            return e.w
        }
    }
    return 0
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    fmt.Println(minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
    fmt.Println(minimumEffortPath([][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}))
    fmt.Println(minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1},
        {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}))
}
