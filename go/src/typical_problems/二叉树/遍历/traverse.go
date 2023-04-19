/**
@author ZhengHao Lou
Date    2022/05/16
*/
package main

func Mid(root *TreeNode) {
	var st []*TreeNode
	for root != nil || len(st) > 0 {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}
		if len(st) > 0 {
			root = st[len(st)-1]
			st = st[:len(st)-1]
			root = root.Right
		}
	}
}
