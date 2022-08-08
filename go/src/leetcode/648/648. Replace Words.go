/**
@author ZhengHao Lou
Date    2022/07/07
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode.cn/problems/replace-words/
*/
const N = 26

type TrieTreeNode struct {
	Value    string
	Children [N]*TrieTreeNode
	IsEnd    bool
}

func NewTrieTree(dictionary []string) *TrieTreeNode {
	root := &TrieTreeNode{}

	var build func(s string)
	build = func(s string) {
		node := root
		for i := range s {
			j := index(s[i])
			if node.Children[j] == nil {
				node.Children[j] = &TrieTreeNode{}
			}
			node = node.Children[j]
		}
		node.IsEnd = true
		node.Value = s
	}

	for _, s := range dictionary {
		build(s)
	}

	return root
}

func (t *TrieTreeNode) Find(s string) string {
	var find func(s string) *TrieTreeNode
	find = func(s string) *TrieTreeNode {
		node := t
		for i := range s {
			j := index(s[i])
			if node.Children[j] == nil {
				return nil
			} else if node.Children[j].IsEnd {
				return node.Children[j]
			} else {
				node = node.Children[j]
			}
		}
		return nil
	}

	x := find(s)
	if x != nil {
		return x.Value
	}
	return s
}

func index(b byte) int {
	return int(b - 'a')
}

func replaceWords(dictionary []string, sentence string) string {
	root := NewTrieTree(dictionary)
	words := strings.Split(sentence, " ")
	var strs []string
	for _, s := range words {
		strs = append(strs, root.Find(s))
	}
	fmt.Println(strs)
	return strings.Join(strs, " ")
}

func main() {
	replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery")
	replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs")
}
