package main

import "fmt"

/**
https://leetcode-cn.com/problems/operations-on-tree/
 */
type TreeNode struct {
	id 			int
	children 	[]*TreeNode
}

type LockingTree struct {
	parent	[]int
	lockMap map[int]int	// node: user
	nodeMap map[int]*TreeNode
	root 	*TreeNode
}

func Constructor(parent []int) LockingTree {
	var mapper = map[int]*TreeNode{}
	for i := range parent {
		if _, ok := mapper[i]; !ok {
			node := &TreeNode{
				id: i,
				children: []*TreeNode{},
			}
			mapper[i] = node
		}

		node := mapper[i]
		pid := parent[i]

		if _, ok := mapper[pid]; !ok {
			pnode := &TreeNode{
				id: pid,
				children: []*TreeNode{},
			}
			mapper[pid] = pnode
		}
		pnode, _ := mapper[pid]
		pnode.children = append(pnode.children, node)
	}

	return LockingTree{
		parent: parent,
		lockMap: map[int]int{},
		nodeMap: mapper,
		root: mapper[0],
	}
}


func (this *LockingTree) Lock(num int, user int) bool {
	if _, ok := this.lockMap[num]; ok {
		return false
	}
	this.lockMap[num] = user
	return true
}


func (this *LockingTree) Unlock(num int, user int) bool {
	if _, ok := this.lockMap[num]; !ok {
		return false
	}
	if user != this.lockMap[num] {
		return false
	}

	delete(this.lockMap, num)
	return true
}


func (this *LockingTree) Upgrade(num int, user int) bool {
	if _, ok := this.lockMap[num]; ok {
		return false
	}

	pid := this.parent[num]
	for pid != -1 {
		if _, ok := this.lockMap[pid]; ok {
			return false
		}
		pid = this.parent[pid]
	}

	var f bool
	root := this.nodeMap[num]
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if _, ok := this.lockMap[node.id]; ok {
			f = true
			delete(this.lockMap, node.id)
		}

		for _, child := range node.children {
			dfs(child)
		}
	}

	dfs(root)

	if !f {
		return false
	}

	this.lockMap[num] = user
	return true
}

func main() {
	obj := Constructor([]int{-1, 0, 0, 1, 1, 2, 2})
	fmt.Println(obj.Lock(2,2))
	fmt.Println(obj.Unlock(2,3))
	fmt.Println(obj.Unlock(2,2))
	fmt.Println(obj.Lock(4,5))
	fmt.Println(obj.Upgrade(0,1))
	fmt.Println(obj.Lock(0,1))
}


/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */
