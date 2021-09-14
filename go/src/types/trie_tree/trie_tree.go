package trie_tree

type TrieTree struct {
	root *TrieNode
}

type TrieNode struct {
	Val string
	Children map[string]*TrieNode
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
			parent.Children[dir] = &TrieNode{
				Val: dir,
				Children: map[string]*TrieNode{},
			}
		}
		parent = parent.Children[dir]
	}
}
