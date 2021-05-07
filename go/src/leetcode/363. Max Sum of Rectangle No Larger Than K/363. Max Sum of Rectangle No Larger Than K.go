package main

import (
	"fmt"
	"math"
)

/**
题目：https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/
题解：https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/solution/ju-xing-qu-yu-bu-chao-guo-k-de-zui-da-sh-70q2/
 */
func maxSumSubmatrix(matrix [][]int, k int) int {
	ans := math.MinInt32
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		sum := make([]int, n)
		for j := i; j < m; j++ {
			for k := 0; k < n; k++ {
				sum[k] += matrix[j][k] // 第j行所有元素和
			}
			b := New()
			b.Put(integer(0))
			s := 0
			for _, v := range sum {
				s += v
				if l := b.LowBound(integer(s - k)); l != nil {
					ans = max(ans, s - l.(integer).IntValue())
				}
				b.Put(integer(s))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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

func (i integer) IntValue() int {
	return int(i)
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


func main() {
	fmt.Println(maxSumSubmatrix([][]int{{1,0,1},{0,-2,3}}, 2))
	fmt.Println(maxSumSubmatrix([][]int{{2,2,-1}}, 3))
}
