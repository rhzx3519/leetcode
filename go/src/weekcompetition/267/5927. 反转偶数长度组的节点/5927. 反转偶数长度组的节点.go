/**
@author ZhengHao Lou
Date    2021/11/14
*/
package main

import "fmt"

type ListNode struct {
	Val int
	Next  *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	for tno := 1; dummy.Next != nil; tno++ {
		var st = []*ListNode{}
		tmp := dummy.Next
		for c := 0; c != tno && tmp != nil; c++ {
			st = append(st, tmp)
			tmp = tmp.Next
		}
		l := len(st)
		if l & 1 == 1 {
			dummy = st[len(st) - 1]
			continue
		}

		for len(st) != 0 {
			r := st[len(st) - 1]
			st = st[:len(st) - 1]
			dummy.Next = r
			dummy = dummy.Next
		}
		dummy.Next = tmp
	}

	print(head)
	return head
}


func main() {
	reverseEvenLengthGroups(build([]int{4,3,0,5,1,2,7,8,6}))		// [4,0,3,5,1,2,7,8,6]
	reverseEvenLengthGroups(build([]int{5,2,6,3,9,1,7,3,8,4}))
	reverseEvenLengthGroups(build([]int{1,1,0,6}))
	reverseEvenLengthGroups(build([]int{2,1}))
	reverseEvenLengthGroups(build([]int{8}))
	reverseEvenLengthGroups(build([]int{0,4,2,1,3}))
}


func build(a []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for i := range a {
		p.Next = &ListNode{Val: a[i]}
		p = p.Next
	}
	return dummy.Next
}

func print(head *ListNode) {
	var a = []int{}
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}
	fmt.Println(a)
}
