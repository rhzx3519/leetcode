package main

import (
	"fmt"
	"sort"
)

func maximumRemovals(s string, p string, removable []int) int {

	l, r := 0, len(removable)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if !isSub(s, p, removable[:mid+1]) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println(l)
	return l
}

func isSub(s, p string, removable []int) bool {
	if len(p) > len(s) - len(removable) {
		return false
	}
	tmp := []int{}
	tmp = append(tmp, removable...)
	sort.Ints(tmp)

	ns, np := len(s), len(p)
	var i, j int
	for i < ns && j < np {
		if contains(tmp, i) {
			i++
			continue
		}
		if s[i] == p[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return j == np
}

func contains(a []int, target int) bool {
	l, r := 0, len(a)-1
	var mid int
	for l <= r {
		mid = l + (r-l)>>1
		if a[mid] == target {
			return true
		} else if a[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	maximumRemovals("abcacb", "ab", []int{3,1,0})
	maximumRemovals("abcbddddd", "abcd", []int{3,2,1,4,5,6})
	maximumRemovals("abcab", "abc", []int{0,1,2,3,4})
	maximumRemovals("qlevcvgzfpryiqlwy", "qlecfqlw", []int{12,5})
}
