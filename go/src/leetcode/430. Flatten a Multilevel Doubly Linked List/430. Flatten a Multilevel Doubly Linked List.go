package main

import "fmt"

type Node struct {
	Val int
	Prev *Node
	Next *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Child != nil {
		sub :=  flatten(root.Child)
		p := sub
		for p.Next != nil {
			p = p.Next
		}
		p.Next = root.Next
		if root.Next != nil {
			root.Next.Prev = p
		}

		root.Next = sub
		sub.Prev =  root

		root.Child = nil
	}
	root.Next = flatten(root.Next)
	return root
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node1.Child = node3
	node1.Next = node2
	node2.Prev = node1

	root := flatten(node1)
	fmt.Println(root)
}
