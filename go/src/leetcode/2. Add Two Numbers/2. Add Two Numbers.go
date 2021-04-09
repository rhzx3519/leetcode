package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	p := dummy

	r, val := 0, 0
	p1, p2 := l1, l2
	for ; p1 != nil || p2 != nil; {
		if p1 != nil && p2 != nil {
			val = p1.Val + p2.Val + r
			p1.Val = val % 10
			r = val / 10
			p.Next = p1
			p1 = p1.Next
			p2 = p2.Next
		} else if p1 != nil {
			val = p1.Val + r
			p1.Val = val % 10
			r = val / 10
			p.Next = p1
			p1 = p1.Next
		} else {
			val = p2.Val + r
			p2.Val = val % 10
			r = val / 10
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if r != 0 {
		p.Next = &ListNode{r, nil}
	}

	return dummy.Next
}
