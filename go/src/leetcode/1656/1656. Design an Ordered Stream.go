/**
@author ZhengHao Lou
Date    2022/08/16
*/
package main

/**
https://leetcode.cn/problems/design-an-ordered-stream/
*/
type OrderedStream struct {
	ptr    int
	stream []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		ptr:    1,
		stream: make([]string, n+1),
	}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	s.stream[idKey] = value
	start := s.ptr
	for s.ptr < len(s.stream) && s.stream[s.ptr] != "" {
		s.ptr++
	}
	return s.stream[start:s.ptr]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
