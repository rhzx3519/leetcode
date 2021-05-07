/**
平衡二叉搜索树的一种简单实现
 */
package balanced_binary_search_tree

import (
	"math/rand"
)

type (
	// T is a empty interface, that is `any` type.
	// since Go is not support generics now(but will coming soon),
	// so we use T to represent any type
	T interface {}

	K Comparable

	Comparable interface {
		CompareTo(o T) int
	}
)

type integer int

func (i integer) CompareTo(o T) int {
	j := o.(integer)
	if j > i {
		return 1
	} else if j < i {
		return 0
	} else { // j == i
		return -1
	}
}

type node struct {
	child 	[2]*node
	priority int
	val 	 Comparable
}

type Tree struct {
	root *node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Put(val Comparable) {
	t.root =  put(t.root, val)

	check(t.root)
}

func (t *Tree) Delete(val Comparable) {
	t.root = delete(t.root, val)

	check(t.root)
}

func (t *Tree) Contains(val Comparable) bool {
	return get(t.root, val) != nil
}

// 返回第一个大于等于val出现的位置
func (t *Tree) LowerBound(val Comparable) Comparable {
	node := lower_bound(t.root, val)
	if node == nil {
		return nil
	}
	return node.val
}

// 返回第一个大于待查数val的位置
func (t *Tree) UpperBound(val Comparable) Comparable {
	node := upper_bound(t.root, val)
	if node == nil {
		return nil
	}
	return node.val
}

// 返回第一个小于val的位置
func (t *Tree) LessBound(val Comparable) Comparable {
	node := less_bound(t.root, val)
	if node == nil {
		return nil
	}
	return node.val
}

func (t *Tree) FindMax() Comparable {
	node := max(t.root)
	if node == nil {
		return nil
	}
	return node.val
}

// ------------------------------------------------------------------------
// *node 's function

// 将d^1子节点换到父节点，将d子节点换到d^1, 将父节点换到d子节点
func (n *node) rotate(d int) *node {
	x := n.child[d^1]
	n.child[d^1] = x.child[d]
	x.child[d] = n
	return x
}

func put(n *node, val Comparable) *node {
	if n == nil {
		return &node{priority: rand.Int(), val: val}
	}
	d := n.val.CompareTo(val)
	if d == -1 {
		return n
	}
	n.child[d] = put(n.child[d], val)
	if n.child[d].priority > n.priority {
		n = n.rotate(d^1)
	}
	return n
}

func delete(n *node, val Comparable) *node {
	if n == nil {
		return nil
	}
	if d := n.val.CompareTo(val); d >= 0 {
		n.child[d] = delete(n.child[d], val)
		return n
	}
	if n.child[1] == nil {
		return n.child[0]
	}
	if n.child[0] == nil {
		return n.child[1]
	}
	d := 0
	if n.child[0].priority > n.child[1].priority {
		d = 1
	}
	n = n.rotate(d)
	n.child[d] = delete(n.child[d], val)
	return n
}

func get(n *node, val Comparable) *node {
	if n == nil {
		return nil
	}
	d := n.val.CompareTo(val)
	if d == -1 {
		return n
	} else {
		return get(n.child[d], val)	
	}
}

func lower_bound(n *node, val Comparable) (lb *node) {
	for n != nil {
		switch c := n.val.CompareTo(val); {
		case c == 0: // val < n.val
			lb = n
			n = n.child[0]
		case c > 0:		// val > n.val
			n = n.child[1]
		default:
			return n
		}
	}
	return
}

func upper_bound(n *node, val Comparable) (lb *node) {
	for n != nil {
		switch c := n.val.CompareTo(val); {
		case c == 0: // val < n.val
			lb = n
			n = n.child[0]
		case c > 0:		// val > n.val
			n = n.child[1]
		default:
			n = n.child[1]
		}
	}
	return
}

// 返回小于target的第一个节点
func less_bound(n *node, val Comparable) (lb *node) {
	for n != nil {
		switch c := n.val.CompareTo(val); {
		case c == 0: // val < n.val
			n = n.child[0]
		case c > 0:		// val > n.val
			lb = n
			n = n.child[1]
		default:
			n = n.child[0]
		}
	}
	return
}

func max(n *node) *node {
	if n == nil {
		return nil
	}
	if n.child[1] == nil {
		return n
	} else {
		return max(n.child[1])
	}
}

// ----------------------------------------------------------------
// check functions

func check(n *node)  {
	if !is_bst(n) {
		panic("check bst failed.")
	}
}

func is_bst(n *node) bool {
	var is func(x *node, min, max Comparable) bool

	is = func(x *node, min, max Comparable) bool {
		if x == nil {
			return true
		}
		if min != nil && (x.val.CompareTo(min) == 1 || x.val.CompareTo(min) == -1) {	// min > x.val -> 1
			return false
		}
		if max != nil && (x.val.CompareTo(max) == 0 || x.val.CompareTo(max) == -1) {
			return false
		}
		return is(x.child[0], min, x.val) && is(x.child[1], x.val, max)
	}
	return is(n, nil, nil)
}

func inorder(n *node, arr *[]Comparable)  {
	if n == nil {
		return
	}
	inorder(n.child[0], arr)
	*arr = append(*arr, n.val)
	inorder(n.child[1], arr)
}












