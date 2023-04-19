/**
@author ZhengHao Lou
Date    2022/06/18
*/
package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		node := &Node{Val: x}
		node.Next = node
		return node
	}

	var head *Node = aNode
	for head.Next != aNode && head.Next.Val >= head.Val {
		head = head.Next
	}
	var pre = head
	var p = head.Next

	fmt.Println(p.Val, head.Val)
	for p.Val < x {
		pre = p
		p = p.Next
		if p == head.Next {
			break
		}
	}
	fmt.Println(p.Val)
	node := &Node{Val: x}
	node.Next = p
	pre.Next = node

	return aNode
}
