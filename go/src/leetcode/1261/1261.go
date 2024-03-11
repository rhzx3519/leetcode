package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type FindElements struct {
    root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        x := node.Val
        if node.Left != nil {
            node.Left.Val = x<<1 + 1
            dfs(node.Left)
        }
        if node.Right != nil {
            node.Right.Val = x<<1 + 2
            dfs(node.Right)
        }
    }
    if root != nil {
        root.Val = 0
    }
    dfs(root)

    return FindElements{
        root: root,
    }
}

func (this *FindElements) Find(target int) bool {
    var st = []int{target}
    for target != 0 {
        if target&1 == 1 {
            target = (target - 1) >> 1
        } else {
            target = (target - 2) >> 1
        }
        st = append(st, target)
    }
    var dfs func(*TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return len(st) == 0
        }
        if len(st) == 0 {
            return false
        }
        if node.Val != st[len(st)-1] {
            return false
        }
        st = st[:len(st)-1]
        if len(st) == 0 {
            return true
        }
        x := st[len(st)-1]
        if x&1 == 1 {
            return dfs(node.Left)
        }
        return dfs(node.Right)
    }

    return dfs(this.root)
}

func main() {
    root := &TreeNode{Val: -1, Right: &TreeNode{Val: -1}}
    s := Constructor(root)
    fmt.Println(s.Find(2))
}
