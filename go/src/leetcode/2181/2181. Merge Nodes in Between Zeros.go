/**
@author ZhengHao Lou
Date    2022/02/21
*/
package main

/**
https://leetcode-cn.com/problems/merge-nodes-in-between-zeros/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for head.Next != nil {
		var s int
		p := head.Next
		for ; p.Val != 0; p = p.Next {
			s += p.Val
		}
		if s > 0 {
			tail.Next = &ListNode{Val: s}
			tail = tail.Next
		}
		head = p
	}
	return dummy.Next
}
