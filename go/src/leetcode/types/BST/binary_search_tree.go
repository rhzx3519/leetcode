package main

import "fmt"

type (
	T interface {}
	K Comparable

	Comparable interface {
		CompareTo(o T) int
	}
)

type integer int

func (i integer) CompareTo(o T) int {
	j := o.(integer)
	if i > j {
		return 1
	} else if i < j {
		return -1
	}
	return 0
}

type node struct {
	key   K
	left  *node
	right *node
}

type BST struct {
	root *node
}

func New() *BST {
	return new(BST)
}

func (b *BST) Contains(k K) bool {
	return get(b.root, k) != nil
}

func (b *BST) Put(k K) {
	b.root = insert(b.root, k)
}

func (b *BST) LowBound(k K) K {
	l := lower_bound(b.root, k)
	if l == nil {
		return nil
	}
	return l.key
}

// ----------------------------------------------------------------------

func get(root *node, k K) *node {
	if root == nil {
		return nil
	}

	cmp := root.key.CompareTo(k)
	if cmp > 0 {
		return get(root.left, k)
	} else if cmp < 0 {
		return get(root.right, k)
	} else {
		return root
	}
}

func delete(root *node, k K) *node {
	if root == nil {
		return nil
	}
	cmp := root.key.CompareTo(k)
	if cmp > 0 {
		return delete(root.left, k)
	} else if cmp < 0 {
		return delete(root.right, k)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			var cur = root.left
			for cur.right != nil { // 寻找左子树的最大节点
				cur = cur.right
			}
			root.key = cur.key
			root.left = delete(root.left, cur.key)
		}
	}
	return root
}

func insert(root *node, k K) *node {
	if root == nil {
		root = &node{
			key: k,
		}
		return root
	}
	cmp := root.key.CompareTo(k)
	if cmp > 0 {
		root.left = insert(root.left, k)
	} else if cmp < 0 {
		root.right = insert(root.right, k)
	}
	return root
}

// 查找二叉树中值>=k的第一个节点
func lower_bound(root *node, k K) (l *node) {
	if root == nil {
		return nil
	}
	for root != nil {
		cmp := root.key.CompareTo(k)
		if cmp >= 0 {
			l = root
			root = root.left
		} else if cmp < 0 {
			root = root.right
		}
	}
	return
}

// 返回第一个大于k的节点
func upper_bound(root *node, k K) (lb *node) {
	if root == nil {
		return nil
	}
	for root != nil {
		switch cmp := root.key.CompareTo(k); {
		case cmp == 0:
			root = root.left
		case cmp > 0:
			lb = root
			root = root.right
		default:
			return root
		}
	}
	return
}

// ---------------------------------------
// tests



func main() {
	bst := New()
	bst.Put(integer(2))
	bst.Put(integer(2))
	bst.Put(integer(5))
	bst.Put(integer(0))

	fmt.Println(bst.LowBound(integer(5)))
	fmt.Println(bst.LowBound(integer(6)))
}