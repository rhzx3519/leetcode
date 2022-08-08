package main

import "fmt"

type DetectSquares struct {
	points map[[2]int]int
	xp 		map[int]map[int]int
}


func Constructor() DetectSquares {
	return DetectSquares{
		points: map[[2]int]int{},
		xp: map[int]map[int]int{},
	}
}


func (this *DetectSquares) Add(point []int)  {
	x, y := point[0], point[1]

	this.points[[2]int{point[0], point[1]}]++
	if _, ok := this.xp[x]; !ok {
		this.xp[x] = make(map[int]int)
	}
	this.xp[x][y]++
}


func (this *DetectSquares) Count(point []int) int {
	var count int
	x, y := point[0], point[1]
	// x-axis
	for yy, c2 := range this.xp[x] {
		if yy == y {
			continue
		}
		l := abs(yy - y)
		{
			// 3rd point
			c3 := this.points[[2]int{x+l, y}]
			c4 := this.points[[2]int{x+l, yy}]
			count += c2 * c3 * c4
		}

		{
			// 3rd point
			c3 := this.points[[2]int{x-l, y}]
			c4 := this.points[[2]int{x-l, yy}]
			count += c2 * c3 * c4
		}
	}
	// y-axis
	fmt.Println(count)
	return count
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}


/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */

func main() {
	var detectSquares = Constructor()
	detectSquares.Add([]int{3, 10})
	detectSquares.Add([]int{11, 2})
	detectSquares.Add([]int{3, 2})
	detectSquares.Count([]int{11, 10}) // 返回 1 。你可以选择：
	//   - 第一个，第二个，和第三个点

	detectSquares.Count([]int{14, 8})
	detectSquares.Add([]int{11, 2})
	detectSquares.Count([]int{11, 10})
}