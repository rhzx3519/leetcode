package main

import "fmt"

func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])

	var rhombus = func(i, j int, c *container) {
		c.put(grid[i][j])
		for k := 1; i-k >= 0 && i+k < m && j-k >= 0 && j+k < n; k++ {
			//lu
			var s int
			for r, c := i, j-k; r >= i-k && c <= j; r, c = r-1, c+1 {
				s += grid[r][c]
			}
			for r, c := i-k+1, j+1; r <= i && c <= j+k; r, c = r+1, c+1 {
				s += grid[r][c]
			}
			for r, c := i+1, j+k-1; r <= i+k && c >= j; r, c = r+1, c-1 {
				s += grid[r][c]
			}
			for r, c := i+k-1, j-1; r > i && c > j-k; r, c = r-1, c-1 {
				s += grid[r][c]
			}
			c.put(s)
		}
	}

	c := newContainer()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rhombus(i, j, &c)
		}
	}

	ans := []int{}
	for i := 0; i < len(c); i++ {
		if c[i] != 0 {
			ans = append(ans, c[i])
		}
	}

	fmt.Println(c)
	return ans
}

type container []int

func newContainer() container {
	return make([]int, 3)
}

func (c container) put(num int) {
	for i := 0; i < len(c); i++ {
		if c[i] == num {
			return
		}
	}
	if num > c[0] {
		copy(c[1:], c[0:2])
		c[0] = num
	} else if num > c[1] {
		c[2] = c[1]
		c[1] = num
	} else if num > c[2] {
		c[2] = num
	}
}

func main() {
	getBiggestThree([][]int{{1,2,3},{4,5,6},{7,8,9}})
	getBiggestThree([][]int{{7,7,7}})
	getBiggestThree([][]int{{3,4,5,1,3},{3,3,4,2,3},{20,30,200,40,10},{1,5,5,4,1},{4,3,2,2,5}})
	getBiggestThree([][]int{
		{20,17,9,13,5,2,9,1,5},
		{14,9,9,9,16,18,3,4,12},
		{18,15,10,20,19,20,15,12,11},
		{19,16,19,18,8,13,15,14,11},
		{4,19,5,2,19,17,7,2,2}})
}