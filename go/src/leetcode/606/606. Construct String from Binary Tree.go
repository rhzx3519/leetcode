/**
@author ZhengHao Lou
Date    2022/03/19
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/construct-string-from-binary-tree/
*/
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func tree2str(root *TreeNode) string {
    if root == nil {
        return "()"
    }
    var bytes = []byte(fmt.Sprintf("%v", root.Val))
    if root.Left == nil && root.Right == nil {
        return string(bytes)
    }

    if root.Left != nil {
        bytes = append(bytes, fmt.Sprintf("(%v)", tree2str(root.Left))...)
    } else {
        bytes = append(bytes, fmt.Sprintf("()")...)
    }
    if root.Right != nil {
        bytes = append(bytes, fmt.Sprintf("(%v)", tree2str(root.Right))...)
    }
    return string(bytes)
}
