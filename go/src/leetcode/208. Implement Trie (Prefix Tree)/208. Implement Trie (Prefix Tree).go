package main

import "fmt"

type node struct {
	val 	rune
	isLeaf 	bool
	child 	[]*node
}

type Trie struct {
	root *node
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: newNode(0),
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	root := this.root
	for _, c := range word {
		i := c - 'a'
		if root.child[i] == nil {
			root.child[i] = newNode(c)
		}
		root = root.child[i]
	}
	root.isLeaf = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this.root
	for _, c := range word {
		i := c - 'a'
		if root.child[i] == nil {
			return false
		}
		root = root.child[i]
	}
	return root != nil && root.isLeaf
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this.root
	for _, c := range prefix {
		i := c - 'a'
		if root.child[i] == nil {
			return false
		}
		root = root.child[i]
	}
	return true
}

// ------------------------------------------------------------

func newNode(val rune) *node {
	return &node{
		val: val,
		child: make([]*node, 26),
	}
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	trie.Insert("apple")

	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.StartsWith("app"))
	fmt.Println(trie.Search("applle"))
	fmt.Println(trie.StartsWith("pp"))
}
