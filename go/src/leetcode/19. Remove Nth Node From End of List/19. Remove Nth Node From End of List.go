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
	p := head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	q := dummy
	for p != nil {
		q = q.Next
		p = p.Next
	}
	q.Next = q.Next.Next
	return dummy.Next
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	fmt.Println(removeNthFromEnd(BuildList([]int{1,2,3,4,5}), 2))
	fmt.Println(removeNthFromEnd(BuildList([]int{1}), 1))
	print(removeNthFromEnd(BuildList([]int{1,2}), 2))
}
