package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	p := dummy
	st := []*ListNode{}
	i := 0
	for ; i < left; i++ {
		head = p
		p = p.Next
	}

	for ; i <= right; i++ {
		st = append(st, p)
		p = p.Next
	}

	for len(st) > 0 {
		top := st[len(st)-1]
		st = st[:len(st)-1]
		head.Next = top
		head = head.Next
	}
	head.Next  = p
	return dummy.Next
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

func main() {
	head := BuildList([]int{1,2,3,4,5})
	reverseBetween(head, 2, 5)
}