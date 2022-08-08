/**
@author ZhengHao Lou
Date    2022/05/11
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return fmt.Sprintf("%v,%v,%v", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	que := strings.Split(data, ",")
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		s := que[0]
		que = que[1:]
		if s == "#" {
			return nil
		}
		i, _ := strconv.Atoi(s)
		root := &TreeNode{Val: i}
		root.Left = dfs()
		root.Right = dfs()
		return root
	}
	return dfs()
}
