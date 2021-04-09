package main

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)

	p := dummy
	for head != nil {
		t := head.Next
		dup := false
		for t != nil && t.Val == head.Val {
			t = t.Next
			dup = true
		}
		if !dup {
			p.Next = head
			p = p.Next
			p.Next = nil
		}
		head = t
	}
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
	head := BuildList([]int{1,1,1,2,3})
	deleteDuplicates(head)
}
