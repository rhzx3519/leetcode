package main

import (
	"fmt"
)

type integer int

func (i integer) CompareTo(o T) int {
	var j = o.(integer)
	if i > j {
		return -1
	} else if i < j {
		return 1
	} else {
		return 0
	}
}

type SeatManager struct {
	n int
	b *BST
}


func Constructor(n int) SeatManager {
	b := New()
	for i := 1; i <= n; i++ {
		b.Put(integer(i))
	}
	return SeatManager{
		n: n,
		b: b,
	}
}


func (this *SeatManager) Reserve() int {
	return int(this.b.DeleteMin().(integer))
}

func (this *SeatManager) Unreserve(seatNumber int)  {
	this.b.Put(integer(seatNumber))
}


type (
	T interface {}
	K Comparable

	Comparable interface {
		CompareTo(o T) int
	}
)

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

func (b *BST) DeleteMin() K {
	minNode := delete_min(b.root)
	if minNode == nil {
		return nil
	}
	delete(b.root, minNode.key)
	return minNode.key
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

func delete_min(root *node) (n *node) {
	if root == nil {
		return nil
	}
	if root.left == nil && root.right == nil {
		return root
	} else if root.left != nil {
		return delete_min(root.left)
	} else if root.right != nil {
		return root
	} else {
		return delete_min(root.left)
	}
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */

func main() {
	n := 5
	seatManager := Constructor(n); // 初始化 SeatManager ，有 5 个座位。
	fmt.Println(seatManager.Reserve())    // 所有座位都可以预约，所以返回最小编号的座位，也就是 1 。
	fmt.Println(seatManager.Reserve())    // 可以预约的座位为 [2,3,4,5] ，返回最小编号的座位，也就是 2 。
	seatManager.Unreserve(2); // 将座位 2 变为可以预约，现在可预约的座位为 [2,3,4,5] 。
	fmt.Println(seatManager.Reserve())   // 可以预约的座位为 [2,3,4,5] ，返回最小编号的座位，也就是 2 。
	fmt.Println(seatManager.Reserve())   // 可以预约的座位为 [3,4,5] ，返回最小编号的座位，也就是 3 。
	fmt.Println(seatManager.Reserve())    // 可以预约的座位为 [4,5] ，返回最小编号的座位，也就是 4 。
	fmt.Println(seatManager.Reserve())    // 唯一可以预约的是座位 5 ，所以返回 5 。
	seatManager.Unreserve(5); // 将座位 5 变为可以预约，现在可预约的座位为 [5] 。
}