package main

import "fmt"

type unit struct {
	value string
	timestamp int
}

type TimeMap struct {
	mapper map[string][]unit
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		mapper: map[string][]unit{},
	}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	if _, ok := this.mapper[key]; !ok {
		this.mapper[key] = []unit{}
	}
	idx := upperBound(this.mapper[key], timestamp)
	this.mapper[key] = append(this.mapper[key], unit{})
	copy(this.mapper[key][idx+1:], this.mapper[key][idx:])
	this.mapper[key][idx] = unit{value: value, timestamp: timestamp}
}


func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.mapper[key]; !ok {
		return ""
	}
	a := this.mapper[key]
	idx := upperBound(a, timestamp)
	if idx == 0 {
		return ""
	}
	fmt.Println(a[idx-1].value)
	return a[idx-1].value
}

func upperBound(a []unit, ts int) int {
	l, r := 0, len(a)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if a[mid].timestamp > ts {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	kv := Constructor()
	kv.Set("foo", "bar", 1); // 存储键 "foo" 和值 "bar" 以及时间戳 timestamp = 1
	kv.Get("foo", 1);  // 输出 "bar"
	kv.Get("foo", 3); // 输出 "bar" 因为在时间戳 3 和时间戳 2 处没有对应 "foo" 的值，所以唯一的值位于时间戳 1 处（即 "bar"）  
	kv.Set("foo", "bar2", 4);
	kv.Get("foo", 4); // 输出 "bar2"
	kv.Get("foo", 5); // 输出 "bar2"


}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
