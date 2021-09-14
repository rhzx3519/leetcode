package main

type node struct {
	name string
	child []*node
}

type ThroneInheritance struct {
	root *node
	death map[string]int
	name2node map[string]*node
}


func Constructor(kingName string) ThroneInheritance {
	root := &node{name: kingName, child: []*node{}}
	return ThroneInheritance{
		root: root,
		death: make(map[string]int),
		name2node: map[string]*node{kingName: root},
	}	
}


func (this *ThroneInheritance) Birth(parentName string, childName string)  {
	p := this.name2node[parentName]
	newNode := &node{name: childName, child: []*node{}}
	this.name2node[childName] = newNode
	p.child = append(p.child, newNode)
}


func (this *ThroneInheritance) Death(name string)  {
	this.death[name]++
}


func (this *ThroneInheritance) GetInheritanceOrder() []string {
	order := []string{}
	this.successor(this.root, &order)
	return order
}

func (this *ThroneInheritance) successor(root *node, order *[]string) {
	if root == nil {
		return
	}
	if _, ok := this.death[root.name]; !ok {
		*order = append(*order, root.name)
	}
	for _, ch := range root.child {
		this.successor(ch, order)
	}
}


/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
