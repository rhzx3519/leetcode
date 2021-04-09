package main

type MyHashMap struct {
	a []int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	hash := MyHashMap{a: make([]int, 1000001)}
	for i := 0; i < len(hash.a); i++ {
		hash.a[i] = -1
	}
	return hash
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	this.a[key] = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.a[key]
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	this.a[key] = -1
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
