package visualize_tree

import (
	"github.com/rhzx3519/leetcode/go/src/leetcode/types"
	"log"
	"math"
)

const (
	ROW_GAP = 100
	COL_GAP = 100
	RADIUS = 30
)

type Pos struct {
	x, y float64
}

func DrawTree(arr []types.T) {
	root := buildTree(arr)
	if root == nil {
		log.Println("empty.")
		return
	}
	// drawer
	drawer := NewDrawer(800, 800)

	var mapper = make(map[*types.TreeNode]Pos)
	mapper[root] = Pos{x: 400, y: 100}
	que := []types.T{root}
	var nodes []*types.TreeNode

	for level := 0; hasNext(&que); level++ {
		nodes = []*types.TreeNode{}

		for sz := len(que); sz > 0; sz-- {
			node := next(&que).(*types.TreeNode)
			nodes = append(nodes, node)
			pos := mapper[node]
			if node.Left != nil {
				que = append(que, node.Left)
				mapper[node.Left] = Pos{x: pos.x-COL_GAP, y: pos.y+ROW_GAP}
			}
			if node.Right != nil {
				que = append(que, node.Right)
				mapper[node.Right] = Pos{x: pos.x+COL_GAP, y: pos.y+ROW_GAP}
			}
		}

		draw(drawer, nodes, mapper)
	}

	drawer.Save("tree.png")
}

func buildTree(que []types.T) *types.TreeNode {
	if !hasNext(&que) {
		return nil
	}
	root := &types.TreeNode{Value: next(&que)}
	lastLevel := []*types.TreeNode{root}
	var tmp []*types.TreeNode
	for hasNext(&que) {
		tmp = make([]*types.TreeNode, 0)
		for _, node := range lastLevel {
			if hasNext(&que) {
				val := next(&que)
				if val != nil {
					node.Left = &types.TreeNode{Value: val, Parent: node}
					tmp = append(tmp, node.Left)
				}
			}
			if hasNext(&que) {
				val := next(&que)
				if val != nil {
					node.Right = &types.TreeNode{Value: val, Parent: node}
					tmp = append(tmp, node.Right)
				}
			}
		}
		lastLevel = tmp
	}
	return root
}

func next(que *[]types.T) (val types.T) {
	val = (*que)[0]
	*que = (*que)[1:]
	return
}

func hasNext(que *[]types.T) bool {
	return len(*que) > 0
}

// 获取树的最大宽度
func width(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + width(root.Left) + width(root.Right)
}

func depth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func draw(drawer *Drawer, nodes []*types.TreeNode, mapper map[*types.TreeNode]Pos) {
	for _, node := range nodes {
		pos := mapper[node]
		drawer.DrawTreeNode(pos.x, pos.y, node)
		if node.Parent == nil {
			continue
		}
		parentPos := mapper[node.Parent]
		length := math.Sqrt((parentPos.x - pos.x)*(parentPos.x - pos.x) + (parentPos.y - pos.y)*(parentPos.y - pos.y))
		sin := (parentPos.y - pos.y) / length
		cos := (parentPos.x - pos.x) / length
		pos.y += RADIUS*sin
		pos.x += RADIUS*cos
		parentPos.y -= RADIUS*sin
		parentPos.x -= RADIUS*cos
		drawer.DrawLine(pos.x, pos.y, parentPos.x, parentPos.y)
	}

}