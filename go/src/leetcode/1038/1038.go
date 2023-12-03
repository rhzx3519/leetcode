package main

/*
*
https://leetcode.cn/problems/binary-search-tree-to-greater-sum-tree/description/?envType=daily-question&envId=2023-12-04
思路：反序中序遍历
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, s int) int {
		if node == nil {
			return s
		}
		right := dfs(node.Right, s)
		node.Val += right
		return dfs(node.Left, node.Val)
	}

	dfs(root, 0)
	return root
}

func main() {
}
