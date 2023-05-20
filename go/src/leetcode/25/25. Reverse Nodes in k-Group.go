package main

/*
*
https://leetcode.cn/problems/reverse-nodes-in-k-group/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	tmp := head
	for i := 1; i < k && tmp != nil; i++ {
		tmp = tmp.Next
	}
	if tmp == nil {
		return head
	}

	cur := reverseKGroup(tmp.Next, k)
	tmp.Next = nil
	newHead := reverse(head)
	head.Next = cur
	return newHead
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverse(head.Next)
	head.Next.Next = head
	return newHead
}
