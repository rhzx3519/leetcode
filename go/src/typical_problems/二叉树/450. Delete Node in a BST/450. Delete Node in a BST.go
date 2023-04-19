package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			// 寻找左子树的最大节点
			tmp := root.Left
			for tmp.Right != nil {
				tmp = tmp.Right
			}
			root.Val = tmp.Val
			root.Left = deleteNode(root.Left, tmp.Val)
		}
	}
	return root
}
