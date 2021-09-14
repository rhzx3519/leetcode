package main

type Node struct {
	Val bool
	IsLeaf bool
	TopLeft *Node
	TopRight *Node
	BottomLeft *Node
	BottomRight *Node
}

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	var node *Node
	if quadTree1.IsLeaf && quadTree2.IsLeaf {
		node = &Node{
			Val: quadTree1.Val || quadTree2.Val,
			IsLeaf: true,
		}
	} else if !quadTree1.IsLeaf && !quadTree2.IsLeaf {
		node = &Node{
			IsLeaf: false,
			TopLeft: intersect(quadTree1.TopLeft, quadTree2.TopLeft),
			TopRight: intersect(quadTree1.TopRight, quadTree2.TopRight),
			BottomLeft: intersect(quadTree1.BottomLeft, quadTree2.BottomLeft),
			BottomRight: intersect(quadTree1.BottomRight, quadTree2.BottomRight),
		}
	} else if quadTree1.IsLeaf {
		if quadTree1.Val {
			node = &Node{
				IsLeaf: true,
				Val: true,
			}
		} else {
			node = &Node{
				IsLeaf: false,
				TopLeft: quadTree2.TopLeft,
				TopRight: quadTree2.TopRight,
				BottomLeft: quadTree2.BottomLeft,
				BottomRight: quadTree2.BottomRight,
			}
		}
	} else {
		if quadTree2.Val {
			node = &Node{
				IsLeaf: true,
				Val: true,
			}
		} else {
			node = &Node{
				IsLeaf: false,
				TopLeft: quadTree1.TopLeft,
				TopRight: quadTree1.TopRight,
				BottomLeft: quadTree1.BottomLeft,
				BottomRight: quadTree1.BottomRight,
			}
		}
	}

	if !node.IsLeaf && node.TopLeft.IsLeaf && node.TopRight.IsLeaf && node.BottomLeft.IsLeaf && node.BottomRight.IsLeaf {
		if node.TopLeft.Val == node.TopRight.Val && node.TopRight.Val == node.BottomLeft.Val && node.BottomLeft.Val == node.BottomRight.Val {
			node.IsLeaf = true
			node.Val = node.TopLeft.Val
			node.TopLeft = nil
			node.TopRight = nil
			node.BottomLeft = nil
			node.BottomRight = nil
		}
	}

	return node
}

