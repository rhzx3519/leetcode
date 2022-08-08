/**
@author ZhengHao Lou
Date    2021/12/23
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/prefix-and-suffix-search/
 */
const N int = 26
type Node struct {
	child []*Node
	words []int
}

func insert(root *Node, i int, w string) {
	p := root
	for _, c := range w {
		j := int(byte(c) - 'a')
		if p.child[j] == nil {
			p.child[j] = &Node{
				child: make([]*Node, N),
				words: []int{},
			}
		}
		p = p.child[j]
		p.words = append(p.words, i)
	}
	p.words = append(p.words, i)
}

func find(root *Node, w string) *Node {
	p := root
	for _, c := range w {
		j := int(byte(c) - 'a')
		if p.child[j] == nil {
			return nil
		}
		p = p.child[j]
	}
	return p
}

type WordFilter struct {
	root *Node
	words []string
}


func Constructor(words []string) WordFilter {
	root := &Node{
		child: make([]*Node, N),
		words: []int{},
	}
	for i, w := range words {
		insert(root, i, w)
	}
	return WordFilter{
		root: root,
		words: words,
	}
}


func (this *WordFilter) F(prefix string, suffix string) int {
	node := find(this.root, prefix)
	if node == nil {
		return -1
	}
	for i := len(node.words) - 1; i >= 0; i-- {
		j := node.words[i]
		if strings.HasSuffix(this.words[j], suffix) {
			return j
		}
	}
	return -1
}

func main() {
	wf := Constructor([]string{"apple"})
	fmt.Println(wf.F("a", "e"))
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
