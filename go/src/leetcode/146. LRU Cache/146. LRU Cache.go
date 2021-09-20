package main

type LRUCache struct {
	capacity int
	cache 	 map[int]int
	priority []int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache: make(map[int]int),
		priority: []int{},
	}
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	i := this.getPriority(key)
	copy(this.priority[1:i+1], this.priority[:i])
	this.priority[0] = key
	return this.cache[key]
}


func (this *LRUCache) Put(key int, value int)  {
	// Put in head
	if _, ok := this.cache[key]; ok {
		this.cache[key] = value
		this.Get(key)
		return
	}
	this.priority = append(this.priority, 0)
	copy(this.priority[1:], this.priority)
	this.priority[0] = key
	this.cache[key] = value
	if len(this.cache) > this.capacity {
		lastKey := this.priority[len(this.priority) - 1]
		delete(this.cache, lastKey)
		this.priority = this.priority[:len(this.priority) - 1]
	}
}

func (this *LRUCache) getPriority(key int) int {
	for i := range this.priority {
		if key == this.priority[i] {
			return i
		}
	}
	return -1
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
