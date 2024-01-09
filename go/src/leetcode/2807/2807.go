package main

/**
https://leetcode.cn/problems/insert-greatest-common-divisors-in-linked-list/?envType=daily-question&envId=2024-01-06
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	for p := head; p.Next != nil; p = p.Next.Next {
		g := gcd(p.Val, p.Next.Val)
		newNode := &ListNode{Val: g}
		newNode.Next = p.Next
		p.Next = newNode
	}
	return head
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
