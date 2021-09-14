package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var pa, pb = headA, headB
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}
		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
	return pa
}
