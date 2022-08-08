/**
@author ZhengHao Lou
Date    2022/07/25
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/complete-binary-tree-inserter/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	nodes []*TreeNode
	root  *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	var tree = CBTInserter{
		nodes: []*TreeNode{{}},
		root:  root,
	}
	if root == nil {
		return tree
	}
	var que = []*TreeNode{root}
	for len(que) != 0 {
		node := que[0]
		que = que[1:]
		if node.Left != nil {
			que = append(que, node.Left)
		}
		if node.Right != nil {
			que = append(que, node.Right)
		}
		tree.nodes = append(tree.nodes, node)
	}
	return tree
}

func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{Val: val}
	this.nodes = append(this.nodes, node)
	index := len(this.nodes) - 1
	var parent = this.nodes[index>>1]
	if index&1 == 0 {
		parent.Left = node
	} else {
		parent.Right = node
	}
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.nodes[1]
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	obj := Constructor(root)
	obj.Insert(3)
	obj.Insert(4)
	fmt.Println(obj.Get_root())
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
