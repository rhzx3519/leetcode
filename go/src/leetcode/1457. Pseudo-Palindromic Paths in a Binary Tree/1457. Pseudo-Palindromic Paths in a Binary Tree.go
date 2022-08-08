/**
@author ZhengHao Lou
Date    2021/12/03
*/
package main

/**
https://leetcode-cn.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

const N = 10
func pseudoPalindromicPaths (root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode, digits [N]int)
	dfs = func(node *TreeNode, digits [10]int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil { // leaf
			digits[node.Val]++
			var count int
			for i := 1; i < N; i++ {
				if digits[i] & 1 == 1 {
					count++
				}
				if count >= 2 {
					break
				}
			}
			if count < 2 {
				ans++
			}
			digits[node.Val]--
			return
		}
		digits[node.Val]++
		dfs(node.Left, digits)
		dfs(node.Right, digits)
		digits[node.Val]--
	}
	dfs(root, [N]int{})

	return ans
}