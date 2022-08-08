/**
@author ZhengHao Lou
Date    2021/11/28
*/
package main

import (
	"fmt"
	"sort"
)

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	mts := map[int]map[int][]int{}
	for _, m := range meetings {
		t := m[2]
		if _, ok := mts[t]; !ok {
			mts[t] = make(map[int][]int)
		}
		a, b := m[0], m[1]
		if _, ok := mts[t][a]; !ok {
			mts[t][a] = []int{}
		}
		if _, ok := mts[t][b]; !ok {
			mts[t][b] = []int{}
		}
		mts[t][a] = append(mts[t][a], b)
		mts[t][b] = append(mts[t][b], a)
	}


	times := []int{}
	for t, _ := range mts {
		times = append(times, t)
	}
	sort.Ints(times)
	known := map[int]bool{}
	known[0] = true
	known[firstPerson] = true
	//fmt.Println(mts)

	var dfs func(idx int, adj map[int][]int)
	dfs = func(idx int, adj map[int][]int) {
		known[idx] = true
		for _, nx := range adj[idx] {
			if known[nx] {
				continue
			}
			known[nx] = true
			dfs(nx, adj)
		}
	}

	for _, t := range times {
		for p := range mts[t] {
			if !known[p] {
				continue
			}
			dfs(p, mts[t])
		}

	}

	keys := []int{}
	for i, _ := range known {
		keys = append(keys, i)
	}
	return keys
}

func main() {
	fmt.Println(findAllPeople(6, [][]int{{1,2,5},{2,3,8},{1,5,10}}, 1))
	fmt.Println(findAllPeople(6, [][]int{{3,1,3},{1,2,2},{0,3,3}}, 3))
	fmt.Println(findAllPeople(5, [][]int{{3,4,2},{1,2,1},{2,3,1}}, 1))
	fmt.Println(findAllPeople(6, [][]int{{0,2,1},{1,3,1},{4,5,1}}, 1))
}

type DisjoinSet struct {
	n int
	parent []int
}

func (d DisjoinSet) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find( d.parent[x])
	}
	return d.parent[x]
}

func (d DisjoinSet) Union(x, y int) bool {
	px, py := d.Find(x), d.Find(y)
	if px != py {
		d.parent[py] = px
		return true
	}
	return false
}

func NewDisjointSet(n int) DisjoinSet {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return DisjoinSet{
		n: n,
		parent: parent,
	}
}