/**
@author ZhengHao Lou
Date    2021/12/06
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/
思路：分别寻找从root到start，从root到dest的路径pa、pb
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {

	var dfs func(node *TreeNode, x int, path *[]byte) bool
	dfs = func(node *TreeNode, x int, path *[]byte) bool {
		if node == nil {
			return false
		}
		if node.Val == x {
			return true
		}
		if dfs(node.Left, x, path) {
			*path = append(*path, 'L')
			return true
		}

		if dfs(node.Right, x, path) {
			*path = append(*path, 'R')
			return true
		}

		return false
	}

	pa, pb := []byte{}, []byte{}
	dfs(root, startValue, &pa)
	dfs(root, destValue, &pb)
	fmt.Println(pa, pb)
	var i, j = len(pa) - 1, len(pb) - 1
	for i >= 0 && j >= 0 && pa[i] == pb[j] {
		pa = pa[:len(pa) - 1]
		pb = pb[:len(pb) - 1]
		i--
		j--
	}

	path := []byte{}
	for i := 0; i < len(pa); i++ {
		path = append(path, 'U')
	}

	for j := len(pb) - 1; j >= 0; j-- {
		path = append(path, pb[j])

	}

	fmt.Println(string(path))
	return string(path)
}


func main() {
	//root := &TreeNode{Val: 5}
	//root.Left = &TreeNode{Val: 1}
	//root.Right = &TreeNode{Val: 2}
	//root.Left.Left = &TreeNode{Val: 3}
	//root.Right.Left = &TreeNode{Val: 6}
	//root.Right.Right = &TreeNode{Val: 4}
	//getDirections(root, 3, 6)

	//root := &TreeNode{Val: 2}
	//root.Left = &TreeNode{Val: 1}
	//getDirections(root, 2, 1)

	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	getDirections(root, 4, 3)

}
