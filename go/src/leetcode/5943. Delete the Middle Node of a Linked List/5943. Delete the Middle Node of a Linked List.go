/**
@author ZhengHao Lou
Date    2021/12/06
*/
package main

/**
https://leetcode-cn.com/problems/delete-the-middle-node-of-a-linked-list/
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	var fast, slow *ListNode
	dummy := &ListNode{}
	dummy.Next = head
	fast = dummy
	slow = dummy
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast == nil {
		pre.Next = slow.Next
	} else {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
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


func main() {
	//head := BuildList([]int{1,3,4,7,1,2,6})
	//head := BuildList([]int{1,3,4})
	//head := BuildList([]int{1,3})
	head := BuildList([]int{1})
	head = deleteMiddle(head)
}
