package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := new(ListNode)
	p := dummy
	st := []*ListNode{}
	for head != nil {
		st = st[:0]
		tmp := head
		cnt := k
		for ; tmp != nil && cnt != 0; tmp, cnt = tmp.Next, cnt-1 {
			st = append(st, tmp)
		}

		if cnt > 0 {
			p.Next = head
			return dummy.Next
		}

		for len(st) != 0 {
			p.Next = st[len(st)-1]
			st = st[:len(st)-1]
			p = p.Next
		}
		p.Next = tmp
		head = tmp
	}
	return dummy.Next
}
