package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	var dfs func(node *ListNode, s int) *ListNode
	dfs = func(node *ListNode, s int) *ListNode {
		if node == nil {
			return nil
		}
		s += node.Val
		if s == 0 {
			return node
		}
		return dfs(node.Next, s)
	}

	dummy := &ListNode{Val: -1}
	dummy.Next = head
	p := dummy
	for p != nil {
		node := dfs(p.Next, 0)
		if node != nil {
			p.Next = node.Next
		} else {
			p = p.Next
		}
	}

	return dummy.Next
}

func main() {
	removeZeroSumSublists(BuildList([]int{0,0}))
	removeZeroSumSublists(BuildList([]int{1,2,3,-3,-2}))
	removeZeroSumSublists(BuildList([]int{1,2,3,-3,4}))
	removeZeroSumSublists(BuildList([]int{1,2,-3,3,1}))
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
