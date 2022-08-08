/**
@author ZhengHao Lou
Date    2022/01/26
*/
package main

/**
https://leetcode-cn.com/problems/detect-squares/
*/
type DetectSquares struct {
	points  [][]int
	xPoints map[int]map[int]int
	yPoints map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		points:  [][]int{},
		xPoints: map[int]map[int]int{},
		yPoints: map[int]map[int]int{},
	}
}

func (this *DetectSquares) Add(point []int) {
	this.points = append(this.points, point)
	x, y := point[0], point[1]
	if _, ok := this.xPoints[x]; !ok {
		this.xPoints[x] = map[int]int{}
	}
	if _, ok := this.yPoints[y]; !ok {
		this.yPoints[y] = map[int]int{}
	}
	this.xPoints[x][y]++
	this.yPoints[y][x]++
}

func (this *DetectSquares) Count(point []int) int {
	var count int
	x, y := point[0], point[1]
	if _, ok := this.xPoints[x]; !ok {
		return 0
	}
	if _, ok := this.yPoints[y]; !ok {
		return 0
	}
	for ty, cy := range this.xPoints[x] {
		if ty-y == 0 {
			continue
		}

		{
			tx := x + ty - y
			cx := this.yPoints[y][tx]
			count += this.xPoints[tx][ty] * cx * cy
		}

		{
			tx := x - ty + y
			cx := this.yPoints[y][tx]
			count += this.xPoints[tx][ty] * cx * cy
		}

	}
	//fmt.Println(count)
	return count
}

func main() {
	//obj := Constructor()
	//obj.Add([]int{3, 10})
	//obj.Add([]int{11, 2})
	//obj.Add([]int{3, 2})
	//obj.Add([]int{3, 2})
	//obj.Count([]int{11, 10})
	//obj.Count([]int{14, 8})
	//obj.Add([]int{11, 2})
	//obj.Count([]int{11, 10})

	// case2
	obj := Constructor()
	obj.Add([]int{419, 351})
	obj.Add([]int{798, 351})
	obj.Add([]int{798, 730})
	obj.Count([]int{419, 730})

}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
