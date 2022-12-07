/**
@author ZhengHao Lou
Date    2022/08/18
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/maximum-equal-frequency/
*/

type Set map[int]int

func (s Set) Add(x int) {
	s[x]++
}

func (s Set) Delete(x int) {
	delete(s, x)
}

func (s Set) Size() int {
	return len(s)
}

func maxEqualFreq(nums []int) int {
	n := len(nums)
	counter := make(map[int]int)
	for i := range nums {
		counter[nums[i]]++
	}
	if len(counter) == 1 {
		return n
	}
	var st = map[int]Set{}
	for k, v := range counter {
		if _, ok := st[v]; !ok {
			st[v] = Set{}
		}
		st[v].Add(k)
	}
	fmt.Println(st)
	for i := n - 1; i >= 0; i-- {
		if len(counter) == 1 {
			return n
		}
		if len(st) == 1 {
			for c := range st {
				if c == 1 {
					return n
				}
			}
		}
		if len(st) == 2 {
			var a1, a2, b1, b2 int
			for c := range st {
				if a1 == 0 {
					a1 = c
					a2 = st[c].Size()
				} else if b1 == 0 {
					b1 = c
					b2 = st[c].Size()
				}
			}
			
			if (a1 == 1 && a2 == 1) || (b1 == 1 && b2 == 1) {
				return n
			}
			
			if a1-b1 == 1 && a2 == 1 {
				return n
			}
			if b1-a1 == 1 && b2 == 1 {
				return n
			}
			//if abs(a1-b1) == 1 && (a2 == 1 || b2 == 1) {
			//	return n
			//}
		}
		x := nums[i]
		k := counter[x]
		st[k].Delete(x)
		if st[k].Size() == 0 {
			delete(st, k)
		}
		counter[x]--
		if counter[x] == 0 {
			delete(counter, x)
			n--
			continue
		}
		k = counter[x]
		if _, ok := st[k]; !ok {
			st[k] = Set{}
		}
		st[k].Add(x)
		n--
	}
	return 0
}

func main() {
	// 7
	fmt.Println(maxEqualFreq([]int{2, 2, 1, 1, 5, 3, 3, 5}))
	// 13
	fmt.Println(maxEqualFreq([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}))
	// 5
	fmt.Println(maxEqualFreq([]int{1, 1, 1, 2, 2, 2}))
	// 8
	fmt.Println(maxEqualFreq([]int{10, 2, 8, 9, 3, 8, 1, 5, 2, 3, 7, 6}))
	//2
	fmt.Println(maxEqualFreq([]int{1, 2}))
	// 5
	fmt.Println(maxEqualFreq([]int{1, 1, 2, 2, 3, 4, 5, 6, 6}))
	// 7
	fmt.Println(maxEqualFreq([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}))
	////fmt.Println(maxEqualFreq([]int{}))
}
