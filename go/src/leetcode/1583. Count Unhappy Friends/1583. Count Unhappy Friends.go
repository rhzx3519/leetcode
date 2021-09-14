package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-unhappy-friends/
 */
func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	d := make(map[int]map[int]int)
	for _, pair := range pairs {
		x, y := pair[0], pair[1]
		if d[x] == nil {
			d[x] = make(map[int]int)
		}

		if d[y] == nil {
			d[y] = make(map[int]int)
		}

		for _, z := range preferences[x] {
			if z == y {
				break
			}
			d[x][z]++ // d[x] 存储比y更亲近的朋友
		}

		for _, z := range preferences[y] {
			if z == x {
				break
			}
			d[y][z]++	// d[y] 存储比x更亲近的朋友
		}
	}

	var count int
	for x, pf := range d {
		for y, _ := range pf {
			if _, ok := d[y]; !ok {
				continue
			}
			if _, ok := d[y][x]; ok {	// y 有比 x更亲近的朋友
				count++
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Println(unhappyFriends(4, [][]int{{1,2,3},{3,2,0},{3,1,0},{1,2,0}}, [][]int{{0,1},{2,3}}))
}
