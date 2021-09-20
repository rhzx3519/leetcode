package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	st := []*ListNode{}
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	for head.Next != nil {
		p := head.Next
		var c = k
		for p != nil && c != 0 {
			st = append(st, p)
			p = p.Next
			c--
		}
		if c != 0 {
			break
		}

		for len(st) > 0 {
			head.Next = st[len(st) - 1]
			st = st[:len(st) - 1]
			head = head.Next
		}
		head.Next = p
	}
	return dummy.Next
}
