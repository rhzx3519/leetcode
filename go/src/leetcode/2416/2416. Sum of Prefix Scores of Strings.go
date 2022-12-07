/**
@author ZhengHao Lou
Date    2022/09/19
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/sum-of-prefix-scores-of-strings/
思路：前缀树
*/
func sumPrefixScores(words []string) []int {
	tree := NewTrieTree()
	for _, w := range words {
		tree.Insert([]byte(w))
	}
	var ans = make([]int, len(words))
	for i, w := range words {
		ans[i] = tree.Traverse([]byte(w))
	}
	fmt.Println(ans)
	return ans
}

func main() {
	sumPrefixScores([]string{"abc", "ab", "bc", "b"})
	sumPrefixScores([]string{"abcd"})
}

const N = 26

type TrieTree struct {
	root *TrieNode
}

type TrieNode struct {
	IsEnd    bool
	Count    int
	Val      byte
	Children []*TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: &TrieNode{
			Val:      '0',
			Children: make([]*TrieNode, N),
		},
	}
}

func (tree *TrieTree) Insert(path []byte) {
	parent := tree.root
	for i := range path {
		j := int(path[i] - 'a')
		if parent.Children[j] == nil {
			parent.Children[j] = &TrieNode{
				Val:      path[i],
				Children: make([]*TrieNode, N),
			}
		}
		parent.Children[j].Count++

		parent = parent.Children[j]
	}
	parent.IsEnd = true
}

func (tree *TrieTree) Traverse(path []byte) (c int) {
	parent := tree.root
	for i := range path {
		j := int(path[i] - 'a')
		c += parent.Children[j].Count
		parent = parent.Children[j]
	}
	return
}
