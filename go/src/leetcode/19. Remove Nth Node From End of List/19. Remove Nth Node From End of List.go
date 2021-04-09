package main

import "fmt"

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	p := dummy
	for ; n >= 0; n-- {
		p = p.Next
	}

	q := dummy
	for ; p != nil; p = p.Next {
		q = q.Next
	}
	if q.Next != nil {
		q.Next = q.Next.Next
	} else {
		q.Next = nil
	}
	return dummy.Next
}

func main() {
	//fmt.Println(removeNthFromEnd(BuildList([]int{1,2,3,4,5}), 2))
	//fmt.Println(removeNthFromEnd(BuildList([]int{1}), 1))
	fmt.Println(removeNthFromEnd(BuildList([]int{1,2}), 2))
}
