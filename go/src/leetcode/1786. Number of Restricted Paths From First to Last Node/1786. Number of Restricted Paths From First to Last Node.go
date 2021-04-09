package main

import (
	"fmt"
	"math"
)

var MOD = int(math.Pow10(9) + 7)
var cache = make(map[int]int)

func countRestrictedPaths(n int, edges [][]int) int {
	cache = make(map[int]int)
	grid := make(map[int][][]int)

	for i := 0; i < len(edges); i++ {
		a, b, c := edges[i][0], edges[i][1], edges[i][2]
		grid[a] = append(grid[a], []int{b, c})
		grid[b] = append(grid[b], []int{a, c})
	}

	dis := dijkstra(n, grid)
	fmt.Println(dis)

	vis := make([]bool, n+1)
	vis[1] = true
	ans := dp(1, n, vis, grid, dis)
	return ans
}

func dijkstra(n int, grid map[int][][]int) []int {
	que := new(PQ)
	dis := make([]int, n+1)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt32;
	}

	que.Push(&entry{key: n, priority: 0})
	for que.Len() > 0 {
		e := que.Pop().(*entry)
		i, c := e.key, e.priority
		if dis[i] < c {
			continue
		}
		for _, adj := range grid[i] {
			ni, nc := adj[0], adj[1]
			if (dis[i] + nc < dis[ni]) {
				dis[ni] = dis[i] + nc
				que.Push(&entry{key: ni, priority: dis[ni]})
			}
		}
	}
	return dis
}

func dp(idx, n int, vis []bool, grid map[int][][]int, dis []int) int {
	if num, ok := cache[idx]; ok {
		return num
	}
	if idx == n {
		return 1
	}
	ans := 0
	for _, adj := range grid[idx] {
		i := adj[0]

		if vis[i] {
			continue
		}
		if dis[i] >= dis[idx] {
			continue
		}
		vis[i] = true
		ans += dp(i, n, vis, grid, dis)%MOD
		vis[i] = false
	}
	cache[idx] = ans%MOD
	return cache[idx]
}

type entry struct {
	key     int
	priority int
	// index 是 entry 在 heap 中的索引号
	// entry 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 entry.priority 一直不变的话，可以删除 index
	index int
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*entry)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1 // for safety
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

func main() {
	fmt.Println(countRestrictedPaths(5, [][]int{{1,2,3},{1,3,3},{2,3,1},{1,4,2},{5,2,2},{3,5,1},{5,4,10}}));
	fmt.Println(countRestrictedPaths(7, [][]int{{1,3,1},{4,1,2},{7,3,4},{2,5,3},{5,6,1},{6,7,2},{7,5,3},{2,6,4}}));
}

