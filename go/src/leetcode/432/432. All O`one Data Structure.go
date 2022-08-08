/**
@author ZhengHao Lou
Date    2021/12/09
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/all-oone-data-structure/
*/
// 十字链表
type CrossListNode struct {
	Key  string
	Val  int
	Pre  *CrossListNode
	Next *CrossListNode
	Up   *CrossListNode
	Down *CrossListNode
}

type AllOne struct {
	dummy *CrossListNode
	tail  *CrossListNode
	hash  map[string]*CrossListNode
	bone  map[int]*CrossListNode
}

func Constructor() AllOne {
	allOne := AllOne{
		dummy: &CrossListNode{Key: "dummy"},
		tail:  &CrossListNode{Key: "tail"},
		hash:  make(map[string]*CrossListNode),
		bone:  make(map[int]*CrossListNode),
	}
	allOne.dummy.Next = allOne.tail
	allOne.tail.Pre = allOne.dummy
	return allOne
}

func (this *AllOne) Inc(key string) {
	if _, ok := this.hash[key]; !ok {
		if _, ok := this.bone[1]; !ok {
			be := &CrossListNode{Val: 1}
			this.bone[1] = be
			be.Next = this.dummy.Next
			be.Next.Pre = be
			this.dummy.Next = be
			be.Pre = this.dummy
		}
		be := this.bone[1]
		node := &CrossListNode{
			Key: key,
			Val: 1,
		}
		node.Up = be
		if be.Down != nil {
			be.Down.Up = node
		}
		node.Down = be.Down
		be.Down = node
		this.hash[key] = node
		return
	}

	node := this.hash[key]
	i := node.Val
	be := this.bone[i]
	ni := i + 1
	// delete from origin position
	if node.Down != nil {
		node.Down.Up = node.Up
	}
	node.Up.Down = node.Down
	// delete bone[i] if it's empty
	if be.Down == nil {
		be.Next.Pre = be.Pre
		be.Pre.Next = be.Next
		this.deleteBone(i)
		be = be.Pre
	}

	if _, ok := this.bone[ni]; !ok { // create bone[ni] if it's not exist
		tmp := &CrossListNode{Val: ni}
		tmp.Next = be.Next
		tmp.Next.Pre = tmp
		tmp.Pre = be
		be.Next = tmp
		this.bone[ni] = tmp
	}

	nextBe := this.bone[ni]
	// insert into new position
	node.Up = nextBe
	if nextBe.Down != nil {
		nextBe.Down.Up = node
	}
	node.Down = nextBe.Down
	nextBe.Down = node
	this.hash[key] = node
	node.Val = ni
}

func (this *AllOne) Dec(key string) {
	if _, ok := this.hash[key]; !ok {
		return
	}

	node := this.hash[key]
	if node.Down != nil {
		node.Down.Up = node.Up
	}
	node.Up.Down = node.Down // delete from origin position
	if node.Val == 1 {       // delete this key
		delete(this.hash, key)
		be := this.bone[1]
		if this.bone[1].Down == nil { // delete bone[1]
			be.Next.Pre = be.Pre
			be.Pre.Next = be.Next
			this.deleteBone(1)
		}
		return
	}
	i := node.Val
	be := this.bone[i]
	if this.bone[i].Down == nil { // delete bone[i]
		be.Next.Pre = be.Pre
		be.Pre.Next = be.Next
		this.deleteBone(i)
		be = be.Pre
	}

	ni := node.Val - 1
	nextBe := this.bone[ni]
	if nextBe == nil {
		nextBe = &CrossListNode{Val: ni}

		if ni < be.Val {
			be.Pre.Next = nextBe
			nextBe.Pre = be.Pre
			nextBe.Next = be
			be.Pre = nextBe
		} else {
			be.Next.Pre = nextBe
			nextBe.Pre = be
			nextBe.Next = be.Next
			be.Next = nextBe
		}

		this.bone[ni] = nextBe
	}
	// insert into new position
	node.Up = nextBe
	if nextBe.Down != nil {
		nextBe.Down.Up = node
	}
	node.Down = nextBe.Down
	nextBe.Down = node
	this.hash[key] = node
	node.Val = ni
}

func (this *AllOne) GetMaxKey() string {
	if this.isEmpty() {
		return ""
	}
	return this.tail.Pre.Down.Key
}

func (this *AllOne) GetMinKey() string {
	if this.isEmpty() {
		return ""
	}
	return this.dummy.Next.Down.Key
}

func (this *AllOne) deleteBone(val int) {
	be := this.bone[val]
	if be.Down != nil {
		return
	}
	be.Next.Pre = be.Pre
	be.Pre.Next = be.Next
	delete(this.bone, val)
}

func (this *AllOne) isEmpty() bool {
	return this.dummy.Next == this.tail
}

func cout(allOne AllOne) {
	fmt.Println("------------------")
	for p := allOne.dummy; p != allOne.tail; p = p.Next {
		fmt.Printf("%v: ", p.Val)
		tmp := []string{}
		for q := p.Down; q != nil; q = q.Down {
			tmp = append(tmp, q.Key)
		}
		fmt.Println(tmp)
	}
}

func main() {
	test2()
	//test3()
}

func test1() {
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("goodbye")
	allOne.Inc("hello")
	cout(allOne)
	allOne.Inc("hello")
	cout(allOne)
	fmt.Println(allOne.GetMaxKey())
	allOne.Inc("leet")
	allOne.Inc("code")
	allOne.Inc("leet")
	allOne.Dec("hello")
	allOne.Inc("leet")
	allOne.Inc("code")
	allOne.Inc("code")
	cout(allOne)
	fmt.Println(allOne.GetMaxKey())
}

//["AllOne","inc","inc","inc","inc","inc","dec","dec","getMaxKey","getMinKey"]
//[[],["a"],["b"],["b"],["b"],["b"],["b"],["b"],[],[]]
func test2() {
	allOne := Constructor()
	allOne.Inc("a")
	allOne.Inc("b")
	allOne.Inc("b")
	allOne.Inc("b")
	allOne.Inc("b")
	cout(allOne)
	allOne.Dec("b")
	cout(allOne)
	allOne.Dec("b")
	cout(allOne)
	fmt.Println(allOne.GetMaxKey())
	fmt.Println(allOne.GetMinKey())
}

//["AllOne","inc","inc","inc","inc","inc","inc","inc","inc","inc","inc","dec","dec","getMaxKey"]
//[[],["hello"],["l"],["l"],["l"],["k"],["k"],["k"],["j"],["j"],["j"],["j"],["k"],[]]
func test3() {
	allOne := Constructor()
	allOne.Inc("l")
	allOne.Inc("l")
	allOne.Inc("l")
	allOne.Inc("k")
	allOne.Inc("k")
	allOne.Inc("k")
	allOne.Inc("j")
	allOne.Inc("j")
	allOne.Inc("j")
	cout(allOne)
	allOne.Dec("j")
	cout(allOne)
	allOne.Dec("k")
	cout(allOne)
	fmt.Println(allOne.GetMaxKey())
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
