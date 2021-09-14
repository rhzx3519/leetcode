package main

import (
	"fmt"
)

var offset = [][]int{{1,0},{-1,0},{0,1},{0,-1},{1,1},{-1,1},{1,-1},{-1,-1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	var step int
	que := NewQueue(100000)
	que.Offer(0)
	grid[0][0] = 1
	for !que.isEmpty() {
		step++
		for sz := que.Size(); sz > 0; sz-- {
			p := que.Poll()
			x, y := p/n, p%n

			if x == n-1 && y == n-1 {
				return step
			}

			for _, dxy := range offset {
				nx := x + dxy[0]
				ny := y + dxy[1]
				if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] != 1 {
					grid[nx][ny] = 1
					que.Offer(nx*n + ny)
				}
			}
		}
	}

	return -1
}

type Queue struct {
	a []int
	cap int
	front, rear int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		a: make([]int, cap),
		cap: cap,
	}
}

func (q *Queue) Offer(x int) {
	if (q.rear + 1) % q.cap == q.front {
		panic("Queue is full.")
	}
	q.a[q.rear] = x
	q.rear = (q.rear + 1) % q.cap
}

func (q *Queue) Poll() (x int) {
	if q.isEmpty() {
		panic("Queue is empty.")
	}
	x = q.a[q.front]
	q.front = (q.front + 1) % q.cap
	return
}

func (q *Queue) isEmpty() bool {
	return q.front == q.rear
}

func (q *Queue) Size() int {
	if q.isEmpty() {
		return 0
	}
	if q.front < q.rear {
		return q.rear - q.front
	}
	return q.rear + q.cap - q.front
}


func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0,1},{1,0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{0,0,0},{1,1,0},{1,1,0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{1,0,0},{1,1,0},{1,1,0}}))
}
