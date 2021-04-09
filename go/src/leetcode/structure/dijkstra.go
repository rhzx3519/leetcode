package structure

import "math"

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