package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	k %= n
	if k == 0 {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	p := dummy
	for i := 0; i != n-k; i++ {
		p = p.Next
	}
	head = p.Next
	p.Next = nil
	p = head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = dummy.Next
	return head
}

type ListNode struct {
	Val int
	Next *ListNode
}
