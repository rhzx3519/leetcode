package main

import "math"
import "github.com/emirpasic/gods/trees/redblacktree"

// https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha
func minimumDistance(points [][]int) int {
    xs := redblacktree.New[int, int]()
    ys := redblacktree.New[int, int]()
    for _, p := range points {
        x, y := p[0]+p[1], p[1]-p[0]
        put(xs, x)
        put(ys, y)
    }

    ans := math.MaxInt
    for _, p := range points {
        x, y := p[0]+p[1], p[1]-p[0]
        remove(xs, x) // 移除一个 x
        remove(ys, y) // 移除一个 y
        ans = min(ans, max(xs.Right().Key-xs.Left().Key, ys.Right().Key-ys.Left().Key))
        put(xs, x)
        put(ys, y)
    }
    return ans
}

func put(t *redblacktree.Tree[int, int], v int) {
    c, _ := t.Get(v)
    t.Put(v, c+1)
}

func remove(t *redblacktree.Tree[int, int], v int) {
    c, _ := t.Get(v)
    if c == 1 {
        t.Remove(v)
    } else {
        t.Put(v, c-1)
    }
}
