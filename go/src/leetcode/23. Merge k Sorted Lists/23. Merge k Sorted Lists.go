package main

import "container/heap"

type ListNode struct {
	Val int
	Next *ListNode
}

func BuildList(a []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, num := range a {
		p.Next = &ListNode{num, nil}
		p = p.Next
	}
	return dummy.Next
}


func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	pq := make(PQ, 0)
	for i := 0; i < n; i++ {
		if lists[i] != nil {
			heap.Push(&pq, lists[i])
		}
	}
	dummy := &ListNode{}
	p := dummy
	for len(pq) != 0 {
		t := heap.Pop(&pq).(*ListNode)
		p.Next = t
		p = p.Next
		if t.Next != nil {
			heap.Push(&pq, t.Next)
		}
	}

	return dummy.Next
}


// PQ implements heap.Interface and holds entries.
type PQ []*ListNode

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*ListNode)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

func main() {

}