/**
给点二叉树中的任意一个节点，输出它的中序遍历后续序列
 */
package main

type Node struct {
	parent, left, right *Node
	val int
}

/**
返回节点root的后继节点
1。 如果右子树不为空，返回右子树的中序第一个节点
2。 否则返回root所处左子树的父节点
 */
func findSucceedNode(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.right != nil {
		root = root.right
		for root.left != nil {
			root = root.left
		}
	} else {
		for root.parent != nil && root != root.parent.left {
			root = root.parent
		}
	}
	return root
}
