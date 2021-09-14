package visualize_tree

import (
	types2 "github.com/rhzx3519/leetcode/go/src/types"
	"log"
	"math"
)

const (
	ROW_GAP = 100
	COL_GAP = 100
	SHRINK_SIZE = 40
	RADIUS = 30
)

type Pos struct {
	x, y float64
}

func DrawTree(arr []types2.T) {
	root := buildTree(arr)
	if root == nil {
		log.Println("empty.")
		return
	}
	// drawer
	drawer := NewDrawer(800, 800)

	var mapper = make(map[*types2.TreeNode]Pos)
	mapper[root] = Pos{x: 400, y: 100}
	que := []types2.T{root}
	var nodes []*types2.TreeNode

	for level := 0; hasNext(&que); level++ {
		nodes = []*types2.TreeNode{}

		for sz := len(que); sz > 0; sz-- {
			node := next(&que).(*types2.TreeNode)
			// calibrate horizontal position
			if len(nodes) > 0 {
				lastNode := nodes[len(nodes)-1]
				lastPos := mapper[lastNode]
				curPos := mapper[node]
				if lastPos.x == curPos.x {
					calibrate(mapper, lastNode, -SHRINK_SIZE)
					calibrate(mapper, node, SHRINK_SIZE)
					//mapper[lastNode] = Pos{x: lastPos.x-SHRINK_SIZE, y: lastPos.y}
					//mapper[node] = Pos{x: curPos.x+SHRINK_SIZE, y: curPos.y}
				}
			}

			nodes = append(nodes, node)

			pos := mapper[node]
			if node.Left != nil {
				que = append(que, node.Left)
				mapper[node.Left] = Pos{x: pos.x - COL_GAP, y: pos.y + ROW_GAP}
			}
			if node.Right != nil {
				que = append(que, node.Right)
				mapper[node.Right] = Pos{x: pos.x + COL_GAP, y: pos.y + ROW_GAP}
			}
		}

		draw(drawer, nodes, mapper)
	}

	drawer.Save("tree.png")
}

func buildTree(que []types2.T) *types2.TreeNode {
	if !hasNext(&que) {
		return nil
	}
	root := &types2.TreeNode{Value: next(&que)}
	lastLevel := []*types2.TreeNode{root}
	var tmp []*types2.TreeNode
	for hasNext(&que) {
		tmp = make([]*types2.TreeNode, 0)
		for _, node := range lastLevel {
			if hasNext(&que) {
				val := next(&que)
				if val != nil {
					node.Left = &types2.TreeNode{Value: val, Parent: node}
					tmp = append(tmp, node.Left)
				}
			}
			if hasNext(&que) {
				val := next(&que)
				if val != nil {
					node.Right = &types2.TreeNode{Value: val, Parent: node}
					tmp = append(tmp, node.Right)
				}
			}
		}
		lastLevel = tmp
	}
	return root
}

func next(que *[]types2.T) (val types2.T) {
	val = (*que)[0]
	*que = (*que)[1:]
	return
}

func hasNext(que *[]types2.T) bool {
	return len(*que) > 0
}

// 获取树的最大宽度
func width(root *types2.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + width(root.Left) + width(root.Right)
}

func depth(root *types2.TreeNode) int {
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

// calibrate horizontal position through tree
func calibrate(mapper map[*types2.TreeNode]Pos, node *types2.TreeNode, size float64)  {
	for size < 0 && node != nil {
		curPos := mapper[node]
		mapper[node] = Pos{x: curPos.x+size, y: curPos.y}
		node = node.Right
	}
	// TODO
	//for node.Parent != nil && ((size < 0 && node == node.Parent.Right) || (size > 0 && node == node.Parent.Left)) {
		curPos := mapper[node]
		mapper[node] = Pos{x: curPos.x+size, y: curPos.y}
		//node = node.Parent
	//}
}

func draw(drawer *Drawer, nodes []*types2.TreeNode, mapper map[*types2.TreeNode]Pos) {
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
		pos.y += RADIUS *sin
		pos.x += RADIUS *cos
		parentPos.y -= RADIUS *sin
		parentPos.x -= RADIUS *cos
		drawer.DrawLine(pos.x, pos.y, parentPos.x, parentPos.y)
	}

}