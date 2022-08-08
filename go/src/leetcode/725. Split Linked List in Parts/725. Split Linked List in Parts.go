package main

type ListNode struct {
	Val int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var n int
	for p := head; p != nil; p = p.Next {
		n++
	}
	var s = n / k
	var r = n % k

	dummy := &ListNode{}
	dummy.Next = head
	ans := []*ListNode{}
	for dummy.Next != nil {
		p := dummy
		var c = s
		if r > 0 {
			c++
			r--
		}
		for p != nil && c > 0 {
			c--
			p = p.Next
		}

		ans = append(ans, dummy.Next)
		dummy.Next = p.Next
		p.Next = nil
	}
	for len(ans) < k {
		ans = append(ans, nil)
	}
	return ans
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3

	splitListToParts(node1, 5)
}
