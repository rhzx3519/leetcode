package tool

type Comparable interface {
	CompareTo(other Comparable) int
}

// PQ implements heap.Interface and holds entries.
type PQ []Comparable

func (pq *PQ) Len() int {
	return len(*pq)
}

func (pq *PQ) Less(i, j int) bool {
	return (*pq)[i].CompareTo((*pq)[j]) <= 0
}

func (pq *PQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PQ) Push(x interface{}) {
	temp := x.(Comparable)
	*pq = append(*pq, temp)
}

func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}