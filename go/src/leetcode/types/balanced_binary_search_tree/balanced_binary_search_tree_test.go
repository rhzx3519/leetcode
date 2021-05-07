package balanced_binary_search_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTree_Put(t *testing.T) {
	tree := NewTree()
	tree.Put(integer(0))
	tree.Put(integer(2))
	tree.Put(integer(1))
	tree.Put(integer(-1))
	tree.Put(integer(-100))
	tree.Put(integer(110))

	arr := []Comparable{}
	inorder(tree.root, &arr)
	log.Println(arr)
}

func TestTree_Delete(t *testing.T) {
	tree := NewTree()
	tree.Put(integer(0))
	tree.Put(integer(2))
	tree.Put(integer(1))
	tree.Put(integer(-1))
	tree.Put(integer(-100))
	tree.Put(integer(110))

	print_tree(tree)

	tree.Delete(integer(0))
	print_tree(tree)

	tree.Delete(integer(-100))
	print_tree(tree)

	tree.Delete(integer(-1000))
	print_tree(tree)

	tree.Delete(integer(-1))
	tree.Delete(integer(1))
	tree.Delete(integer(2))
	tree.Delete(integer(110))

	print_tree(tree)
	assert.Nil(t, tree.root)
}

func TestTree_DeleteNotExist(t *testing.T) {
	tree := NewTree()
	tree.Put(integer(0))

	tree.Delete(integer(-1))
	print_tree(tree)
}

func TestTree_Contains(t *testing.T) {
	tree := NewTree()
	tree.Put(integer(0))
	tree.Put(integer(2))
	tree.Put(integer(1))
	tree.Put(integer(-1))
	tree.Put(integer(-100))
	tree.Put(integer(110))

	assert.True(t, tree.Contains(integer(0)))
	assert.True(t, tree.Contains(integer(2)))
	assert.True(t, tree.Contains(integer(1)))
	assert.True(t, tree.Contains(integer(-1)))
	assert.True(t, tree.Contains(integer(-100)))
	assert.True(t, tree.Contains(integer(110)))

	assert.False(t, tree.Contains(integer(9)))
	assert.False(t, tree.Contains(integer(12)))
}

func TestTree_rotate(t *testing.T) {
	const N = 1<<10
	tree := NewTree()
	for i := 0; i < N; i++ {
		tree.Put(integer(i))
	}

	print_tree_by_lv(tree)
}

func TestTree_LowerBound(t *testing.T) {
	tree := NewTree()
	tree.Put(integer(0))
	tree.Put(integer(2))
	tree.Put(integer(1))
	tree.Put(integer(-1))
	tree.Put(integer(-100))
	tree.Put(integer(110))

	assert.Equal(t, integer(1), tree.LowerBound(integer(1)))
	assert.Equal(t, integer(110), tree.LowerBound(integer(5)))
	assert.Equal(t, integer(110), tree.LowerBound(integer(10)))
	assert.Equal(t, integer(-1), tree.LowerBound(integer(-2)))
	assert.Equal(t, integer(-100), tree.LowerBound(integer(-101)))
	assert.Equal(t, integer(-1), tree.LowerBound(integer(-99)))

	assert.Nil(t, tree.LowerBound(integer(111)))
}

func TestTree_UpperBound(t *testing.T) {
	tree := NewTree()
	tree.Put(integer(0))
	tree.Put(integer(2))
	tree.Put(integer(1))
	tree.Put(integer(-1))
	tree.Put(integer(-100))
	tree.Put(integer(110))

	assert.Equal(t, integer(2), tree.UpperBound(integer(1)))
	assert.Equal(t, integer(110), tree.UpperBound(integer(2)))
	assert.Equal(t, integer(0), tree.UpperBound(integer(-1)))
	assert.Equal(t, integer(-1), tree.UpperBound(integer(-100)))
	assert.Equal(t, integer(-100), tree.UpperBound(integer(-101)))

	assert.Nil(t, tree.UpperBound(integer(110)))
}

func TestTree_LessBound(t *testing.T) {
	tree := NewTree()
	tree.Put(integer(0))
	tree.Put(integer(2))
	tree.Put(integer(1))
	tree.Put(integer(-1))
	tree.Put(integer(-100))
	tree.Put(integer(110))

	assert.Equal(t, integer(0), tree.LessBound(integer(1)))
	assert.Equal(t, integer(-1), tree.LessBound(integer(0)))
	assert.Equal(t, integer(-100), tree.LessBound(integer(-1)))
	assert.Equal(t, integer(-100), tree.LessBound(integer(-2)))
	assert.Equal(t, nil, tree.LessBound(integer(-100)))
	assert.Equal(t, integer(110), tree.LessBound(integer(1000)))
	assert.Equal(t, integer(1), tree.LessBound(integer(2)))
	assert.Equal(t, integer(2), tree.LessBound(integer(3)))
	assert.Equal(t, integer(2), tree.LessBound(integer(100)))

}

func print_tree(tree *Tree)  {
	arr := []Comparable{}
	inorder(tree.root, &arr)
	log.Println(arr)
}

func print_tree_by_lv(tree *Tree) {
	if tree.root == nil {
		return
	}
	que := []*node{tree.root}
	lv := 0
	for len(que) > 0 {
		for sz := len(que); sz > 0; sz-- {
			root := que[0]
			fmt.Print(root.val, " ")
			que = que[1:]
			for _, cld := range root.child {
				if cld != nil {
					que = append(que, cld)
				}
			}
		}
		lv++
		fmt.Println()
	}

	fmt.Println("\n total level: ", lv)
}

