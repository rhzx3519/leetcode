package main

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var mapper = make(map[*Node]*Node)
	var copy func(node *Node) *Node
	copy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if _, ok := mapper[node]; ok {
			return mapper[node]

		}
		mapper[node] = &Node{
			Val: node.Val,
		}
		c := mapper[node]
		c.Random = copy(node.Random)
		return c
	}

	var dummy = &Node{}
	var c = dummy
	for head != nil {
		c.Next = copy(head)
		head = head.Next
		c = c.Next
	}
	return dummy.Next
}
