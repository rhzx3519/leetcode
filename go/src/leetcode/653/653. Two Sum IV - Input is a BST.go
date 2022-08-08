/**
@author ZhengHao Lou
Date    2022/03/21
*/
package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
    counter := make(map[int]int)
    var dfs func(root *TreeNode) bool
    dfs = func(root *TreeNode) bool {
        if root == nil {
            return false
        }
        if counter[k-root.Val] != 0 {
            return true
        }
        counter[root.Val]++
        return dfs(root.Left) || dfs(root.Right)
    }

    return dfs(root)
}
