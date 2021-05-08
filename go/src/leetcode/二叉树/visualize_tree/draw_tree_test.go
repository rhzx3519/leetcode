package visualize_tree

import (
	"github.com/rhzx3519/leetcode/go/src/leetcode/types"
	"testing"
)

const S = 800

func TestDrawer_drawCirle(t *testing.T) {

	drawer := NewDrawer(S, S)
	drawer.DrawCirle(200, 200, 10)
	drawer.Save("circle.png")
}

func TestDrawer_drawText(t *testing.T) {
	drawer := NewDrawer(S, S)
	drawer.DrawText(200, 200, "Hello World")
	drawer.Save("text.png")
}

func TestDrawer_drawLine(t *testing.T) {
	drawer := NewDrawer(S, S)
	drawer.DrawLine(100, 100, 300, 100)
	drawer.Save("line.png")
}

func TestDrawer_drawTreeNode(t *testing.T) {
	node := &types.TreeNode{
		Value: 123,
	}
	drawer := NewDrawer(S, S)
	drawer.DrawTreeNode(200, 200, node)
	drawer.Save("node.png")
}