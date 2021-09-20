package main

import "fmt"

/**
https://leetcode-cn.com/problems/minimize-hamming-distance-after-swap-operations/solution/bing-cha-ji-yao-shao-wei-dong-dian-nao-j-feis/
思路：并查集 + 哈希表统计不同元素个数
allowedSwaps需要用并查集来表示，e.g. (0,1), (1,2) -> (0,2) 也可以交换的，所以需要用并查集来维护这个关系,

 */
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	parent = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for _, as := range allowedSwaps {
		union(as[0], as[1])
	}

	var s = make(map[int]map[int]int)
	var t = make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		pi := find(i)
		if _, ok := s[pi]; !ok {
			s[pi] = make(map[int]int)
		}
		if _, ok := t[pi]; !ok {
			t[pi] = make(map[int]int)
		}

		s[pi][source[i]]++
		t[pi][target[i]]++
		// s[pi]存储了可以通过交换放到 pi 的source中的所有元素；t[pi]同理
		// 两个集合s[pi], t[pi]的最小汉明距就是两个集合不同元素的个数, 具体求法是：
		// 1. 遍历第一个集合s[pi], 如果元素出现在t[pi]集合，则将该元素从t[pi]中删除；
		// 2. 否则 汉明距 += 1
	}

	var diff int
	for i := 0; i < n; i++ {
		if _, ok := s[i]; !ok {
			continue
		}

		// e.g. [1,2,3], [1,2,2] -> {1:1, 2:1, 3:1}, {1:1, 2:2}
		for num, cs := range s[i] {
			ct := t[i][num]
			t[i][num] = max(0, ct - cs)	// 统计s[i], p[i] 两个集合中的不同元素数量
		}

		for _, c := range t[i] {
			diff += c
		}
	}

	return diff
}

var parent []int

func find(x int) int{
	if parent[x] != x {
		return find(parent[x])
	}
	return parent[x]
}

func union(x, y int) {
	px := find(x)
	py := find(y)
	if px != py {
		parent[py] = px
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(minimumHammingDistance([]int{1,2,3,4}, []int{2,1,4,5}, [][]int{{0,1}, {2,3}}))	// 1
	fmt.Println(minimumHammingDistance([]int{1,2,3,4}, []int{1,3,2,4}, [][]int{}))		// 2
	fmt.Println(minimumHammingDistance([]int{5,1,2,4,3}, []int{1,5,4,2,3}, [][]int{{0,4},{4,2},{1,3},{1,4}})) // 0
	fmt.Println(minimumHammingDistance([]int{2,3,1}, []int{1,2,2}, [][]int{{0,2},{1,2}}))	// 1
}