package tool


type Comparable interface {
	CompareTo(other Comparable) int
}

// PQ implements heap.Interface and holds entries.
type PQ []Comparable

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].CompareTo(pq[j]) <= 0
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(Comparable)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}