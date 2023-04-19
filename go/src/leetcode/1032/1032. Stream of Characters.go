package main

import "fmt"

/*
*
https://leetcode.cn/problems/stream-of-characters/
*/
type StreamChecker struct {
	words []string
	que   []byte
	root  *TrieNode
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{
		words: words,
		root:  &TrieNode{child: [N]*TrieNode{}},
	}
	for _, w := range words {
		sc.root.insert(w)
	}

	return sc
}

func (this *StreamChecker) Query(letter byte) bool {
	this.que = append(this.que, letter)
	return this.root.query(string(this.que))
}

const N = 26

type TrieNode struct {
	child [N]*TrieNode
	end   bool
}

func (t *TrieNode) insert(w string) {
	p := t
	for i := len(w) - 1; i >= 0; i-- {
		index := int(w[i] - 'a')
		if p.child[index] == nil {
			p.child[index] = &TrieNode{child: [N]*TrieNode{}}
		}
		p = p.child[index]
	}
	p.end = true
}

func (t *TrieNode) query(w string) bool {
	p := t
	n := len(w)
	for i := n - 1; i >= min(0, 200-n); i-- {
		index := int(w[i] - 'a')
		if p.child[index] == nil {
			return false
		}
		p = p.child[index]
		if p.end {
			return true
		}
	}
	return p.end
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sc := Constructor([]string{"cd", "f", "kl"})
	sc.Query('a')
	sc.Query('b')
	sc.Query('c')
	fmt.Println(sc.Query('d'))
}
