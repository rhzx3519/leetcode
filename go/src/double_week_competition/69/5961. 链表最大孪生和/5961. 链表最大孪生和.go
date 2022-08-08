/**
@author ZhengHao Lou
Date    2022/01/08
*/
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow := head
	fast := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	st := []*ListNode{}
	for slow != nil {
		st = append(st, slow)
		slow = slow.Next
	}
	var ans int
	for len(st) != 0 {
		t := head.Val + st[len(st)-1].Val
		ans = max(ans, t)
		st = st[:len(st)-1]
		head = head.Next
	}

	fmt.Println(ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	pairSum(head)
}
