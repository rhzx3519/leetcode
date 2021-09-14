package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Node struct {
	Val int
	Left *Node
	Right *Node
	Parent *Node
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var t *Node

	var traverse func(node *TreeNode) *Node
	traverse = func(node *TreeNode) *Node {
		if node == nil {
			return nil
		}
		root := &Node{
			Val: node.Val,
			Left: traverse(node.Left),
			Right: traverse(node.Right),
		}
		if target == node {
			t = root
		}
		if root.Left != nil {
			root.Left.Parent = root
		}
		if root.Right != nil {
			root.Right.Parent = root
		}
		return root
	}

	traverse(root)

	vis := map[*Node]bool{}
	var ans = []int{}
	que := []*Node{t}
	vis[t] = true
	for ; len(que) > 0 && k >= 0; k-- {
		for sz := len(que); sz > 0; sz-- {
			t := que[0]
			que = que[1:]
			ans = append(ans, t.Val)
			if t.Left != nil && !vis[t.Left] {
				vis[t.Left] = true
				que = append(que, t.Left)
			}
			if t.Right != nil && !vis[t.Right] {
				vis[t.Right] = true
				que = append(que, t.Right)
			}
			if t.Parent != nil && !vis[t.Parent] {
				vis[t.Parent] = true
				que = append(que, t.Parent)
			}
		}
		ans = []int{}
	}
	return ans
}
