/*
*
@author ZhengHao Lou
Date    2022/11/28
*/
package main

/*
*
https://leetcode.cn/problems/remove-nodes-from-linked-list/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	var st []*ListNode
	for ; head != nil; head = head.Next {
		for len(st) != 0 && st[len(st)-1].Val < head.Val {
			st = st[:len(st)-1]
		}
		st = append(st, head)

	}
	dummy := &ListNode{}
	p := dummy
	for i := range st {
		p.Next = st[i]
		p = p.Next
	}
	return dummy.Next
}
