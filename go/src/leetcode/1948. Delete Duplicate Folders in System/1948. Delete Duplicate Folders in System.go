package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
https://leetcode-cn.com/problems/delete-duplicate-folders-in-system/
思路：前缀树

 */
var mapper map[string][]*TrieNode

func deleteDuplicateFolder(paths [][]string) [][]string {
	mapper = map[string][]*TrieNode{}

	trie := NewTrieTree()
	for _, path := range paths {
		trie.Insert(path)
	}
	serialize(trie.root)

	for _, nodes := range mapper {
		if len(nodes) <= 1 {
			continue
		}
		for _, node := range nodes {
			trie.root = delete(trie.root, node.Parent)
		}
	}
	var ans = [][]string{}

	var dfs func(root *TrieNode, arr []string)
	dfs = func(root *TrieNode, arr []string) {
		if root == nil {
			return
		}
		if len(arr) > 0 {
			ans = append(ans, append([]string{}, arr...))
		}

		for _, child := range root.Children {
			if child == nil {
				continue
			}
			arr = append(arr, child.Val)
			dfs(child, arr)
			arr = arr[:len(arr) - 1]
		}
	}

	dfs(trie.root, []string{})

	fmt.Println(ans)
	return ans
}

// 序列化TrieTree
func serialize(root *TrieNode) string {
	if root == nil {
		return "#"
	}
	if root.Str != "" {
		return root.Str
	}
	arr := []string{root.Val}

	keys := []string{}
	for key, _ := range root.Children {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		s := serialize(root.Children[key])
		arr = append(arr, s)
	}
	s := strings.Join(arr, ".")
	root.Str = s
	mapper[s] = append(mapper[s], root)
	return s
}

func delete(root, node *TrieNode) *TrieNode {
	if node == nil || root == nil {
		return root
	}
	if root == node {
		root = nil
		return nil
	}
	for key, _ := range root.Children {
		root.Children[key] = delete(root.Children[key], node)
	}
	return root
}

func main() {
	//deleteDuplicateFolder([][]string{{"a"},{"c"},{"d"},{"a","b"},{"c","b"},{"d","a"}})
	// [["d"],["d","a"]]
	//deleteDuplicateFolder([][]string{{"a"},{"c"},{"a","b"},{"c","b"},{"a","b","x"},{"a","b","x","y"},{"w"},{"w","y"}})
	// [["c"],["c","b"],["a"],["a","b"]]
	//deleteDuplicateFolder([][]string{{"a"},{"a","x"},{"a","x","y"},{"a","z"},{"b"},{"b","x"},{"b","x","y"},{"b","z"}})
	// []
	deleteDuplicateFolder([][]string{{"a"},{"a","x"},{"a","x","y"},{"a","z"},{"b"},{"b","x"},{"b","x","y"},{"b","z"},{"b","w"}})
	// [["b"],["b","w"],["b","z"],["a"],["a","z"]]
}

type TrieTree struct {
	root *TrieNode
}

type TrieNode struct {
	Val string
	Parent *TrieNode
	Children map[string]*TrieNode
	Str string
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: &TrieNode{
			Val: "",
			Children: map[string]*TrieNode{},
		},
	}
}

func (tree *TrieTree) Insert(path []string) {
	parent := tree.root
	for _, dir := range path {
		if _, ok := parent.Children[dir]; !ok {
			child := &TrieNode{
				Val: dir,
				Parent: parent,
				Children: map[string]*TrieNode{},
			}
			parent.Children[dir] = child
		}
		parent = parent.Children[dir]
	}
}